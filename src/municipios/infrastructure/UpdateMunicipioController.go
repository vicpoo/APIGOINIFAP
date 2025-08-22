// UpdateMunicipioController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/application"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/domain/entities"
)

type UpdateMunicipioController struct {
	updateUseCase *application.UpdateMunicipioUseCase
}

func NewUpdateMunicipioController(updateUseCase *application.UpdateMunicipioUseCase) *UpdateMunicipioController {
	return &UpdateMunicipioController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateMunicipioController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var municipioRequest struct {
		ClaveEstado    int32  `json:"clave_estado"`
		ClaveMunicipio int32  `json:"clave_municipio"`
		Nombre         string `json:"nombre"`
	}

	if err := c.ShouldBindJSON(&municipioRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	municipio := entities.NewMunicipio(
		municipioRequest.ClaveEstado,
		municipioRequest.ClaveMunicipio,
		municipioRequest.Nombre,
	)
	municipio.ID = int32(id)

	updatedMunicipio, err := ctrl.updateUseCase.Run(municipio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el municipio",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedMunicipio)
}