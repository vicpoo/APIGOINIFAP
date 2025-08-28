//GetClasificacionResultadosByMunicipioController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/application"
)

type GetClasificacionResultadosByMunicipioController struct {
	getByMunicipioUseCase *application.GetClasificacionResultadosByMunicipioUseCase
}

func NewGetClasificacionResultadosByMunicipioController(getByMunicipioUseCase *application.GetClasificacionResultadosByMunicipioUseCase) *GetClasificacionResultadosByMunicipioController {
	return &GetClasificacionResultadosByMunicipioController{
		getByMunicipioUseCase: getByMunicipioUseCase,
	}
}

func (ctrl *GetClasificacionResultadosByMunicipioController) Run(c *gin.Context) {
	municipioIDParam := c.Param("municipio_id")
	municipioID, err := strconv.Atoi(municipioIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de municipio inválido",
			"error":   err.Error(),
		})
		return
	}

	clasificaciones, err := ctrl.getByMunicipioUseCase.Run(int32(municipioID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al obtener clasificaciones del municipio",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, clasificaciones)
}

