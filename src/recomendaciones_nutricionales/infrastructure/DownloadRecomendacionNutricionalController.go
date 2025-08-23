// DownloadRecomendacionNutricionalController.go
package infrastructure

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/application"
)

type DownloadRecomendacionNutricionalController struct {
	downloadUseCase *application.DownloadRecomendacionNutricionalUseCase
	fileUploader    *FileUploader
}

func NewDownloadRecomendacionNutricionalController(downloadUseCase *application.DownloadRecomendacionNutricionalUseCase) *DownloadRecomendacionNutricionalController {
	return &DownloadRecomendacionNutricionalController{
		downloadUseCase: downloadUseCase,
		fileUploader:    NewFileUploader(),
	}
}

// RunByID descarga un PDF por ID de recomendación
func (ctrl *DownloadRecomendacionNutricionalController) RunByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	recomendacion, err := ctrl.downloadUseCase.RunByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Recomendación no encontrada",
			"error":   err.Error(),
		})
		return
	}

	// Verificar que el archivo existe
	if _, err := os.Stat(recomendacion.RutaPDF); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Archivo PDF no encontrado",
			"error":   "El archivo no existe en el servidor",
		})
		return
	}

	// Configurar headers para descarga
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+recomendacion.NombrePDF)
	c.Header("Content-Type", "application/pdf")
	c.Header("Cache-Control", "no-cache")

	// Servir el archivo
	c.File(recomendacion.RutaPDF)
}

// RunByMunicipio obtiene todas las recomendaciones de un municipio
func (ctrl *DownloadRecomendacionNutricionalController) RunByMunicipio(c *gin.Context) {
	municipioIDParam := c.Param("municipio_id")
	municipioID, err := strconv.Atoi(municipioIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de municipio inválido",
			"error":   err.Error(),
		})
		return
	}

	recomendaciones, err := ctrl.downloadUseCase.RunByMunicipio(int32(municipioID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al obtener recomendaciones del municipio",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recomendaciones)
}

// RunDownloadByMunicipio descarga todos los PDFs de un municipio como archivo ZIP
func (ctrl *DownloadRecomendacionNutricionalController) RunDownloadByMunicipio(c *gin.Context) {
	municipioIDParam := c.Param("municipio_id")
	municipioID, err := strconv.Atoi(municipioIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de municipio inválido",
			"error":   err.Error(),
		})
		return
	}

	recomendaciones, err := ctrl.downloadUseCase.RunByMunicipio(int32(municipioID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al obtener recomendaciones del municipio",
			"error":   err.Error(),
		})
		return
	}

	if len(recomendaciones) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No se encontraron recomendaciones para este municipio",
		})
		return
	}

	// Crear archivo ZIP temporal
	zipFilename := "recomendaciones_municipio_" + municipioIDParam + ".zip"
	zipPath := filepath.Join(os.TempDir(), zipFilename)

	err = ctrl.fileUploader.CreateZipFromPDFs(recomendaciones, zipPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al crear archivo ZIP",
			"error":   err.Error(),
		})
		return
	}

	// Limpiar archivo temporal después de enviar
	defer os.Remove(zipPath)

	// Configurar headers para descarga ZIP
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+zipFilename)
	c.Header("Content-Type", "application/zip")
	c.Header("Cache-Control", "no-cache")

	// Servir el archivo ZIP
	c.File(zipPath)
}