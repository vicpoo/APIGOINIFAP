//GetAllRolesController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/rol/application"
)

type GetAllRolesController struct {
	getAllUseCase *application.GetAllRolesUseCase
}

func NewGetAllRolesController(getAllUseCase *application.GetAllRolesUseCase) *GetAllRolesController {
	return &GetAllRolesController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllRolesController) Run(c *gin.Context) {
	roles, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los roles",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, roles)
}
