//DeleteRolController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/rol/application"
)

type DeleteRolController struct {
	deleteUseCase *application.DeleteRolUseCase
}

func NewDeleteRolController(deleteUseCase *application.DeleteRolUseCase) *DeleteRolController {
	return &DeleteRolController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteRolController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el rol",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Rol eliminado exitosamente",
	})
}
