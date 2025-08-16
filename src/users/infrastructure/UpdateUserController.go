// UpdateUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/users/application"
	"github.com/vicpoo/APIGOINIFAP/src/users/domain/entities"
)

type UpdateUserController struct {
	updateUseCase *application.UpdateUserUseCase
}

func NewUpdateUserController(updateUseCase *application.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{updateUseCase: updateUseCase}
}

func (ctrl *UpdateUserController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	var request struct {
		Nombre           string `json:"nombre"`
		Apellido         string `json:"apellido"`
		Correo           string `json:"correo"`
		NumeroTelefonico string `json:"numero_telefonico"`
		Password         string `json:"password"`
		RolIDFK          int32  `json:"rol_id_FK"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos", "error": err.Error()})
		return
	}

	user := entities.NewUser(request.Nombre, request.Apellido, request.Correo, request.NumeroTelefonico, request.Password, request.RolIDFK)
	user.SetIDUser(int32(id))

	updatedUser, err := ctrl.updateUseCase.Run(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo actualizar el usuario", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
