// GetAllRecomendacionesNutricionalesController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/application"
)

type GetAllRecomendacionesNutricionalesController struct {
	getAllUseCase *application.GetAllRecomendacionesNutricionalesUseCase
}

func NewGetAllRecomendacionesNutricionalesController(getAllUseCase *application.GetAllRecomendacionesNutricionalesUseCase) *GetAllRecomendacionesNutricionalesController {
	return &GetAllRecomendacionesNutricionalesController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllRecomendacionesNutricionalesController) Run(c *gin.Context) {
	recomendaciones, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las recomendaciones nutricionales",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recomendaciones)
}