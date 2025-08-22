// GetAllMunicipiosController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/application"
)

type GetAllMunicipiosController struct {
	getAllUseCase *application.GetAllMunicipiosUseCase
}

func NewGetAllMunicipiosController(getAllUseCase *application.GetAllMunicipiosUseCase) *GetAllMunicipiosController {
	return &GetAllMunicipiosController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllMunicipiosController) Run(c *gin.Context) {
	municipios, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los municipios",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, municipios)
}