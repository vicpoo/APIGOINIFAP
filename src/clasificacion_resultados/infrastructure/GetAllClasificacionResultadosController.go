//GetAllClasificacionResultadosController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/application"
)

type GetAllClasificacionResultadosController struct {
	getAllUseCase *application.GetAllClasificacionResultadosUseCase
}

func NewGetAllClasificacionResultadosController(getAllUseCase *application.GetAllClasificacionResultadosUseCase) *GetAllClasificacionResultadosController {
	return &GetAllClasificacionResultadosController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllClasificacionResultadosController) Run(c *gin.Context) {
	clasificaciones, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las clasificaciones",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, clasificaciones)
}

