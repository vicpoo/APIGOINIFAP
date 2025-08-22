// UpdateRecomendacionNutricionalController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/application"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type UpdateRecomendacionNutricionalController struct {
	updateUseCase *application.UpdateRecomendacionNutricionalUseCase
}

func NewUpdateRecomendacionNutricionalController(updateUseCase *application.UpdateRecomendacionNutricionalUseCase) *UpdateRecomendacionNutricionalController {
	return &UpdateRecomendacionNutricionalController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateRecomendacionNutricionalController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

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
	recomendacion.ID = int32(id)

	updatedRecomendacion, err := ctrl.updateUseCase.Run(recomendacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la recomendación nutricional",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedRecomendacion)
}