// GetMunicipioByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/application"
)

type GetMunicipioByIdController struct {
	getByIdUseCase *application.GetMunicipioByIdUseCase
}

func NewGetMunicipioByIdController(getByIdUseCase *application.GetMunicipioByIdUseCase) *GetMunicipioByIdController {
	return &GetMunicipioByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetMunicipioByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	municipio, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el municipio",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, municipio)
}