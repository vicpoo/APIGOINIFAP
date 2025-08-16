// rol_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type RolRouter struct {
	engine *gin.Engine
}

func NewRolRouter(engine *gin.Engine) *RolRouter {
	return &RolRouter{
		engine: engine,
	}
}

func (router *RolRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, getAllController := InitRolDependencies()

	// Grupo de rutas para roles
	rolGroup := router.engine.Group("/roles")
	{
		rolGroup.POST("/", createController.Run)
		rolGroup.GET("/:id", getByIdController.Run)
		rolGroup.PUT("/:id", updateController.Run)
		rolGroup.DELETE("/:id", deleteController.Run)
		rolGroup.GET("/", getAllController.Run)
	}
}