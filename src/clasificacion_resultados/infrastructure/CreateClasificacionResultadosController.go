//CreateClasificacionResultadosController.go
package infrastructure

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/application"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain/entities"
)

type CreateClasificacionResultadosController struct {
	createUseCase *application.CreateClasificacionResultadosUseCase
}

func NewCreateClasificacionResultadosController(createUseCase *application.CreateClasificacionResultadosUseCase) *CreateClasificacionResultadosController {
	return &CreateClasificacionResultadosController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateClasificacionResultadosController) Run(c *gin.Context) {
	var request struct {
		MunicipioID        int32  `json:"municipio_id_FK" binding:"required"`
		AnalisisTipo       string `json:"analisis_tipo" binding:"required"`
		FechaAnalisis      string `json:"fecha_analisis" binding:"required"`
		ResultadoGeneral   string `json:"resultado_general" binding:"required"`
		NutrientesCriticos string `json:"nutrientes_criticos"`
		Comentario         string `json:"comentario"`
		UserID             int32  `json:"user_id_FK" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Parsear fecha
	fechaAnalisis, err := time.Parse("2006-01-02", request.FechaAnalisis)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fecha inválido. Use YYYY-MM-DD",
			"error":   err.Error(),
		})
		return
	}

	clasificacion := entities.NewClasificacionResultados(
		request.MunicipioID,
		request.AnalisisTipo,
		fechaAnalisis,
		request.ResultadoGeneral,
		request.NutrientesCriticos,
		request.Comentario,
		"", // Imagen vacía inicialmente
		request.UserID,
	)

	createdClasificacion, err := ctrl.createUseCase.Run(clasificacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la clasificación de resultados",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdClasificacion)
}