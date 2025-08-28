//DeleteClasificacionResultadosController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/application"
)

type DeleteClasificacionResultadosController struct {
	deleteUseCase *application.DeleteClasificacionResultadosUseCase
}

func NewDeleteClasificacionResultadosController(deleteUseCase *application.DeleteClasificacionResultadosUseCase) *DeleteClasificacionResultadosController {
	return &DeleteClasificacionResultadosController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteClasificacionResultadosController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	// Verificar que existe antes de eliminar
	_, err = ctrl.deleteUseCase.GetById(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Clasificación no encontrada",
			"error":   err.Error(),
		})
		return
	}

	// Eliminar el registro
	err = ctrl.deleteUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar la clasificación",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Clasificación eliminada exitosamente",
	})
}