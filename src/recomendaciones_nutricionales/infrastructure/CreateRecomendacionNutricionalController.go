// CreateRecomendacionNutricionalController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/application"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type CreateRecomendacionNutricionalController struct {
	createUseCase *application.CreateRecomendacionNutricionalUseCase
}

func NewCreateRecomendacionNutricionalController(createUseCase *application.CreateRecomendacionNutricionalUseCase) *CreateRecomendacionNutricionalController {
	return &CreateRecomendacionNutricionalController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateRecomendacionNutricionalController) Run(c *gin.Context) {
	var recomendacionRequest struct {
		MunicipioID int32  `json:"municipio_id_FK"`
		NombrePDF   string `json:"nombre_pdf"`
		RutaPDF     string `json:"ruta_pdf"`
		UserID      int32  `json:"user_id_FK"`
	}

	if err := c.ShouldBindJSON(&recomendacionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	recomendacion := entities.NewRecomendacionNutricional(
		recomendacionRequest.MunicipioID,
		recomendacionRequest.NombrePDF,
		recomendacionRequest.RutaPDF,
		recomendacionRequest.UserID,
	)

	createdRecomendacion, err := ctrl.createUseCase.Run(recomendacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la recomendación nutricional",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdRecomendacion)
}