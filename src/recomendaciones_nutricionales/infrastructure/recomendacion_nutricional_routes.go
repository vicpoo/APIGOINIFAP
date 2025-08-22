// recomendacion_nutricional_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type RecomendacionNutricionalRouter struct {
	engine *gin.Engine
}

func NewRecomendacionNutricionalRouter(engine *gin.Engine) *RecomendacionNutricionalRouter {
	return &RecomendacionNutricionalRouter{
		engine: engine,
	}
}

func (router *RecomendacionNutricionalRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, getAllController := InitRecomendacionNutricionalDependencies()

	// Grupo de rutas para recomendaciones nutricionales
	recomendacionGroup := router.engine.Group("/recomendaciones-nutricionales")
	{
		recomendacionGroup.POST("/", createController.Run)
		recomendacionGroup.GET("/:id", getByIdController.Run)
		recomendacionGroup.PUT("/:id", updateController.Run)
		recomendacionGroup.DELETE("/:id", deleteController.Run)
		recomendacionGroup.GET("/", getAllController.Run)
	}
}