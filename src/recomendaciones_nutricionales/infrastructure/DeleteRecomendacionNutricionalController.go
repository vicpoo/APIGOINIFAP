// DeleteRecomendacionNutricionalController.go
package infrastructure

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/application"
)

type DeleteRecomendacionNutricionalController struct {
	deleteUseCase *application.DeleteRecomendacionNutricionalUseCase
	fileUploader  *FileUploader
}

func NewDeleteRecomendacionNutricionalController(deleteUseCase *application.DeleteRecomendacionNutricionalUseCase) *DeleteRecomendacionNutricionalController {
	return &DeleteRecomendacionNutricionalController{
		deleteUseCase: deleteUseCase,
		fileUploader:  NewFileUploader(),
	}
}

func (ctrl *DeleteRecomendacionNutricionalController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	// Obtener la recomendación primero usando el método GetById del caso de uso
	recomendacion, err := ctrl.deleteUseCase.GetById(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Recomendación no encontrada",
			"error":   err.Error(),
		})
		return
	}

	// Eliminar el registro de la base de datos
	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar la recomendación nutricional",
			"error":   errDelete.Error(),
		})
		return
	}

	// Eliminar el archivo físico del PDF
	if err := ctrl.fileUploader.DeleteFile("pdfs", recomendacion.NombrePDF); err != nil {
		// Logear el error pero no fallar la operación completa
		fmt.Printf("Advertencia: No se pudo eliminar el archivo físico %s: %v\n", recomendacion.NombrePDF, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Recomendación nutricional eliminada exitosamente",
	})
}