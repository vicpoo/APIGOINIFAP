//GetClasificacionResultadosByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/application"
)

type GetClasificacionResultadosByIdController struct {
	getByIdUseCase *application.GetClasificacionResultadosByIdUseCase
}

func NewGetClasificacionResultadosByIdController(getByIdUseCase *application.GetClasificacionResultadosByIdUseCase) *GetClasificacionResultadosByIdController {
	return &GetClasificacionResultadosByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetClasificacionResultadosByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	clasificacion, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Clasificación no encontrada",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, clasificacion)
}

