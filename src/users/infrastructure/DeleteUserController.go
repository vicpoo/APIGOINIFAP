//DeleteUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/users/application"
)

type DeleteUserController struct {
	deleteUseCase *application.DeleteUserUseCase
}

func NewDeleteUserController(deleteUseCase *application.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{deleteUseCase: deleteUseCase}
}

func (ctrl *DeleteUserController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	err = ctrl.deleteUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo eliminar el usuario", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Usuario eliminado exitosamente"})
}