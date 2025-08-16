package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/rol/application"
)

type GetRolByIdController struct {
	getByIdUseCase *application.GetRolByIdUseCase
}

func NewGetRolByIdController(getByIdUseCase *application.GetRolByIdUseCase) *GetRolByIdController {
	return &GetRolByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetRolByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	rol, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el rol",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rol)
}
