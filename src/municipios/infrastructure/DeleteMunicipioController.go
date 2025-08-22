// DeleteMunicipioController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/application"
)

type DeleteMunicipioController struct {
	deleteUseCase *application.DeleteMunicipioUseCase
}

func NewDeleteMunicipioController(deleteUseCase *application.DeleteMunicipioUseCase) *DeleteMunicipioController {
	return &DeleteMunicipioController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteMunicipioController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el municipio",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Municipio eliminado exitosamente",
	})
}