package infrastructure

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type FileUploader struct {
	BaseUploadPath string
}

func NewFileUploader() *FileUploader {
	return &FileUploader{
		BaseUploadPath: "uploads",
	}
}

// UploadImage maneja la subida de imágenes para clasificación de resultados
func (fu *FileUploader) UploadImage(c *gin.Context, fileField string) (string, string, error) {
	allowedTypes := []string{"image/jpeg", "image/png", "image/gif", "image/jpg", "image/webp", "image/bmp"}
	return fu.uploadFile(c, fileField, "images/clasificacion", allowedTypes...)
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
		return "", "", fmt.Errorf("tipo de archivo no permitido: %s. Tipos permitidos: %v", mimeType, allowedTypes)
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
	uniqueFilename := fmt.Sprintf("clasificacion_%s%s", timestamp, ext)
	
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

// GetFileInfo obtiene información del archivo
func (fu *FileUploader) GetFileInfo(filePath string) (os.FileInfo, error) {
	return os.Stat(filePath)
}

// FileExists verifica si un archivo existe
func (fu *FileUploader) FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// ServeImage sirve una imagen específica
func (fu *FileUploader) ServeImage(c *gin.Context, folder, filename string) {
	filePath := filepath.Join(fu.BaseUploadPath, folder, filename)
	
	if !fu.FileExists(filePath) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Imagen no encontrada"})
		return
	}

	// Determinar content type basado en extensión
	ext := strings.ToLower(filepath.Ext(filename))
	var contentType string
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	case ".webp":
		contentType = "image/webp"
	case ".bmp":
		contentType = "image/bmp"
	default:
		contentType = "application/octet-stream"
	}

	c.Header("Content-Type", contentType)
	c.File(filePath)
}