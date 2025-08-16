// user_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	engine *gin.Engine
}

func NewUserRouter(engine *gin.Engine) *UserRouter {
	return &UserRouter{
		engine: engine,
	}
}

func (router *UserRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, getAllController, loginController := InitUserDependencies()

	// Grupo de rutas para usuarios
	userGroup := router.engine.Group("/users")
	{
		userGroup.POST("/", createController.Run)         // Crear usuario
		userGroup.GET("/:id", getByIdController.Run)      // Obtener usuario por ID
		userGroup.PUT("/:id", updateController.Run)       // Actualizar usuario
		userGroup.DELETE("/:id", deleteController.Run)    // Eliminar usuario
		userGroup.GET("/", getAllController.Run)          // Obtener todos los usuarios
		userGroup.POST("/login", loginController.Run)     // Login
	}
}
