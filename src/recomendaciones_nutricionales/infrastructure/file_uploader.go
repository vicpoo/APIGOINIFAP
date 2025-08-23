// file_uploader.go
package infrastructure

import (
	"archive/zip"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type FileUploader struct {
	BaseUploadPath string
}

func NewFileUploader() *FileUploader {
	return &FileUploader{
		BaseUploadPath: "uploads",
	}
}

// UploadPDF maneja la subida de archivos PDF
func (fu *FileUploader) UploadPDF(c *gin.Context, fileField string) (string, string, error) {
	return fu.uploadFile(c, fileField, "pdfs", "application/pdf")
}

// UploadImage maneja la subida de imágenes
func (fu *FileUploader) UploadImage(c *gin.Context, fileField string) (string, string, error) {
	allowedTypes := []string{"image/jpeg", "image/png", "image/gif", "image/jpg"}
	return fu.uploadFile(c, fileField, "images", allowedTypes...)
}

// uploadFile maneja la subida genérica de archivos
func (fu *FileUploader) uploadFile(c *gin.Context, fileField, folder string, allowedTypes ...string) (string, string, error) {
	file, header, err := c.Request.FormFile(fileField)
	if err != nil {
		return "", "", fmt.Errorf("error al obtener el archivo: %v", err)
	}
	defer file.Close()

	// Verificar tipo MIME
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", "", fmt.Errorf("error al leer el archivo: %v", err)
	}
	file.Seek(0, 0) // Resetear el reader

	mimeType := http.DetectContentType(buffer)
	
	// Verificar si el tipo MIME está permitido
	allowed := false
	for _, allowedType := range allowedTypes {
		if strings.HasPrefix(mimeType, allowedType) {
			allowed = true
			break
		}
	}

	if !allowed {
		return "", "", fmt.Errorf("tipo de archivo no permitido: %s", mimeType)
	}

	// Generar nombre único para el archivo
	ext := filepath.Ext(header.Filename)
	if ext == "" {
		// Si no tiene extensión, intentar determinar por MIME type
		exts, _ := mime.ExtensionsByType(mimeType)
		if len(exts) > 0 {
			ext = exts[0]
		} else {
			ext = ".bin"
		}
	}

	timestamp := time.Now().Format("20060102150405")
	uniqueFilename := fmt.Sprintf("%s_%s%s", strings.TrimSuffix(header.Filename, ext), timestamp, ext)
	
	// Crear ruta completa
	uploadPath := filepath.Join(fu.BaseUploadPath, folder)
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		return "", "", fmt.Errorf("error al crear directorio: %v", err)
	}

	filePath := filepath.Join(uploadPath, uniqueFilename)

	// Crear archivo destino
	out, err := os.Create(filePath)
	if err != nil {
		return "", "", fmt.Errorf("error al crear archivo: %v", err)
	}
	defer out.Close()

	// Copiar contenido
	_, err = io.Copy(out, file)
	if err != nil {
		return "", "", fmt.Errorf("error al guardar archivo: %v", err)
	}

	return uniqueFilename, filePath, nil
}

// DeleteFile elimina un archivo
func (fu *FileUploader) DeleteFile(folder, filename string) error {
	filePath := filepath.Join(fu.BaseUploadPath, folder, filename)
	return os.Remove(filePath)
}

// GetFileURL genera la URL para acceder al archivo
func (fu *FileUploader) GetFileURL(folder, filename string) string {
	return fmt.Sprintf("/uploads/%s/%s", folder, filename)
}

// CreateZipFromPDFs crea un archivo ZIP con múltiples PDFs
func (fu *FileUploader) CreateZipFromPDFs(recomendaciones []entities.RecomendacionNutricional, zipPath string) error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("error al crear archivo ZIP: %v", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, recomendacion := range recomendaciones {
		// Verificar que el archivo existe
		if _, err := os.Stat(recomendacion.RutaPDF); os.IsNotExist(err) {
			continue // Saltar archivos que no existen
		}

		// Abrir archivo PDF
		pdfFile, err := os.Open(recomendacion.RutaPDF)
		if err != nil {
			continue // Saltar archivos que no se pueden abrir
		}
		defer pdfFile.Close()

		// Crear entrada en el ZIP
		zipEntry, err := zipWriter.Create(recomendacion.NombrePDF)
		if err != nil {
			pdfFile.Close()
			continue
		}

		// Copiar contenido del PDF al ZIP
		_, err = io.Copy(zipEntry, pdfFile)
		if err != nil {
			continue
		}
	}

	return nil
}

// GetFileInfo obtiene información del archivo
func (fu *FileUploader) GetFileInfo(filePath string) (os.FileInfo, error) {
	return os.Stat(filePath)
}

// FileExists verifica si un archivo existe
func (fu *FileUploader) FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}