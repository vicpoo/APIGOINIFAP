// municipio_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type MunicipioRouter struct {
	engine *gin.Engine
}

func NewMunicipioRouter(engine *gin.Engine) *MunicipioRouter {
	return &MunicipioRouter{
		engine: engine,
	}
}

func (router *MunicipioRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, getAllController := InitMunicipioDependencies()

	// Grupo de rutas para municipios
	municipioGroup := router.engine.Group("/municipios")
	{
		municipioGroup.POST("/", createController.Run)
		municipioGroup.GET("/:id", getByIdController.Run)
		municipioGroup.PUT("/:id", updateController.Run)
		municipioGroup.DELETE("/:id", deleteController.Run)
		municipioGroup.GET("/", getAllController.Run)
	}
}