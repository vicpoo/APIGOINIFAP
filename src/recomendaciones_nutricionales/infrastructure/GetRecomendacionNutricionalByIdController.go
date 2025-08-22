// GetRecomendacionNutricionalByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/application"
)

type GetRecomendacionNutricionalByIdController struct {
	getByIdUseCase *application.GetRecomendacionNutricionalByIdUseCase
}

func NewGetRecomendacionNutricionalByIdController(getByIdUseCase *application.GetRecomendacionNutricionalByIdUseCase) *GetRecomendacionNutricionalByIdController {
	return &GetRecomendacionNutricionalByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetRecomendacionNutricionalByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	recomendacion, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la recomendación nutricional",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recomendacion)
}