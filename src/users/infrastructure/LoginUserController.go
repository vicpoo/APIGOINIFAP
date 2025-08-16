// LoginUserController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/users/application"
)

type LoginUserController struct {
	loginUseCase *application.LoginUserUseCase
}

func NewLoginUserController(loginUseCase *application.LoginUserUseCase) *LoginUserController {
	return &LoginUserController{loginUseCase: loginUseCase}
}

func (ctrl *LoginUserController) Run(c *gin.Context) {
	var request struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	user, err := ctrl.loginUseCase.Run(request.Correo, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Credenciales inválidas",
			"error":   err.Error(),
		})
		return
	}

	// Retornar usuario + token por separado
	c.JSON(http.StatusOK, gin.H{
		"id_user":           user.IDUser,
		"nombre":            user.Nombre,
		"apellido":          user.Apellido,
		"correo":            user.Correo,
		"numero_telefonico": user.NumeroTelefonico,
		"rol_id_FK":         user.RolIDFK,
		"token":             user.Password, // Aquí va el token generado
	})
}
