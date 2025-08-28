//UploadImageClasificacionResultadosController.go
package infrastructure

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/application"
)

type UploadImageClasificacionResultadosController struct {
	uploadImageUseCase *application.UploadImageClasificacionResultadosUseCase
}

func NewUploadImageClasificacionResultadosController(uploadImageUseCase *application.UploadImageClasificacionResultadosUseCase) *UploadImageClasificacionResultadosController {
	return &UploadImageClasificacionResultadosController{
		uploadImageUseCase: uploadImageUseCase,
	}
}

func (ctrl *UploadImageClasificacionResultadosController) Run(c *gin.Context) {
	// Obtener ID de la clasificación
	idParam := c.Param("id")
	clasificacionID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	// Obtener archivo de imagen
	file, err := c.FormFile("imagen")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error al obtener archivo de imagen",
			"error":   err.Error(),
		})
		return
	}

	// Validar tipo de archivo
	if !isImageFile(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "El archivo debe ser una imagen (jpg, jpeg, png, gif)",
		})
		return
	}

	// Directorio de upload (deberías configurar esto según tu estructura)
	uploadDir := "./uploads/images"

	// Subir imagen
	updatedClasificacion, err := ctrl.uploadImageUseCase.Run(
		int32(clasificacionID),
		file,
		uploadDir,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al subir imagen",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedClasificacion)
}

// isImageFile valida que el archivo sea una imagen
func isImageFile(filename string) bool {
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".webp": true,
	}
	
	ext := strings.ToLower(filepath.Ext(filename))
	return allowedExtensions[ext]
}