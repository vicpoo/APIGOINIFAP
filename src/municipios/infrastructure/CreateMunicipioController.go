// CreateMunicipioController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/application"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/domain/entities"
)

type CreateMunicipioController struct {
	createUseCase *application.CreateMunicipioUseCase
}

func NewCreateMunicipioController(createUseCase *application.CreateMunicipioUseCase) *CreateMunicipioController {
	return &CreateMunicipioController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateMunicipioController) Run(c *gin.Context) {
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

	createdMunicipio, err := ctrl.createUseCase.Run(municipio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el municipio",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdMunicipio)
}