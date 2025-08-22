// DeleteRecomendacionNutricionalController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/application"
)

type DeleteRecomendacionNutricionalController struct {
	deleteUseCase *application.DeleteRecomendacionNutricionalUseCase
}

func NewDeleteRecomendacionNutricionalController(deleteUseCase *application.DeleteRecomendacionNutricionalUseCase) *DeleteRecomendacionNutricionalController {
	return &DeleteRecomendacionNutricionalController{
		deleteUseCase: deleteUseCase,
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

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar la recomendación nutricional",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Recomendación nutricional eliminada exitosamente",
	})
}