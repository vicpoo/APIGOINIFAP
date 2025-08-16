// GetUserByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/users/application"
)

type GetUserByIdController struct {
	getByIdUseCase *application.GetUserByIdUseCase
}

func NewGetUserByIdController(getByIdUseCase *application.GetUserByIdUseCase) *GetUserByIdController {
	return &GetUserByIdController{getByIdUseCase: getByIdUseCase}
}

func (ctrl *GetUserByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	user, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo obtener el usuario", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
