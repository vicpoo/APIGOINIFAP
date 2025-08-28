//UpdateClasificacionResultadosController.go
package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/application"
)

type UpdateClasificacionResultadosController struct {
	updateUseCase *application.UpdateClasificacionResultadosUseCase
}

func NewUpdateClasificacionResultadosController(updateUseCase *application.UpdateClasificacionResultadosUseCase) *UpdateClasificacionResultadosController {
	return &UpdateClasificacionResultadosController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateClasificacionResultadosController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	// Obtener clasificación existente
	existingClasificacion, err := ctrl.updateUseCase.GetById(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Clasificación no encontrada",
			"error":   err.Error(),
		})
		return
	}

	var request struct {
		MunicipioID        int32  `json:"municipio_id_FK"`
		AnalisisTipo       string `json:"analisis_tipo"`
		FechaAnalisis      string `json:"fecha_analisis"`
		ResultadoGeneral   string `json:"resultado_general"`
		NutrientesCriticos string `json:"nutrientes_criticos"`
		Comentario         string `json:"comentario"`
		UserID             int32  `json:"user_id_FK"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Parsear fecha si se proporciona
	var fechaAnalisis time.Time
	if request.FechaAnalisis != "" {
		fechaAnalisis, err = time.Parse("2006-01-02", request.FechaAnalisis)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Formato de fecha inválido. Use YYYY-MM-DD",
				"error":   err.Error(),
			})
			return
		}
	} else {
		fechaAnalisis = existingClasificacion.FechaAnalisis
	}

	// Actualizar campos
	if request.MunicipioID != 0 {
		existingClasificacion.MunicipioID = request.MunicipioID
	}
	if request.AnalisisTipo != "" {
		existingClasificacion.AnalisisTipo = request.AnalisisTipo
	}
	existingClasificacion.FechaAnalisis = fechaAnalisis
	if request.ResultadoGeneral != "" {
		existingClasificacion.ResultadoGeneral = request.ResultadoGeneral
	}
	if request.NutrientesCriticos != "" {
		existingClasificacion.NutrientesCriticos = request.NutrientesCriticos
	}
	if request.Comentario != "" {
		existingClasificacion.Comentario = request.Comentario
	}
	if request.UserID != 0 {
		existingClasificacion.UserID = request.UserID
	}

	updatedClasificacion, err := ctrl.updateUseCase.Run(existingClasificacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la clasificación",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedClasificacion)
}