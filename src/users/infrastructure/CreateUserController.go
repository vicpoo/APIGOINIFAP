//CreateUserController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/users/application"
	"github.com/vicpoo/APIGOINIFAP/src/users/domain/entities"
)

type CreateUserController struct {
	createUseCase *application.CreateUserUseCase
}

func NewCreateUserController(createUseCase *application.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{createUseCase: createUseCase}
}

func (ctrl *CreateUserController) Run(c *gin.Context) {
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

	createdUser, err := ctrl.createUseCase.Run(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo crear el usuario", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
