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
	createController, getByIdController, updateController, deleteController, getAllController, downloadController := InitRecomendacionNutricionalDependencies()

	// Grupo de rutas para recomendaciones nutricionales
	recomendacionGroup := router.engine.Group("/recomendaciones-nutricionales")
	{
		recomendacionGroup.POST("/", createController.Run)
		recomendacionGroup.GET("/:id", getByIdController.Run)
		recomendacionGroup.PUT("/:id", updateController.Run)
		recomendacionGroup.DELETE("/:id", deleteController.Run)
		recomendacionGroup.GET("/", getAllController.Run)
		
		// Nuevas rutas para descargas
		recomendacionGroup.GET("/:id/download", downloadController.RunByID)
		recomendacionGroup.GET("/municipio/:municipio_id", downloadController.RunByMunicipio)
		recomendacionGroup.GET("/municipio/:municipio_id/download-zip", downloadController.RunDownloadByMunicipio)
	}
}