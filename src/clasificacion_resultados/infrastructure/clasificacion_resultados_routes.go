package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type ClasificacionResultadosRouter struct {
	engine *gin.Engine
}

func NewClasificacionResultadosRouter(engine *gin.Engine) *ClasificacionResultadosRouter {
	return &ClasificacionResultadosRouter{
		engine: engine,
	}
}

func (router *ClasificacionResultadosRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, 
	deleteController, getAllController, getByMunicipioController, 
	uploadImageController := InitClasificacionResultadosDependencies()

	// File uploader para servir imágenes
	fileUploader := NewFileUploader()

	// Grupo de rutas para clasificación de resultados
	clasificacionGroup := router.engine.Group("/clasificacion-resultados")
	{
		// CRUD básico
		clasificacionGroup.POST("/", createController.Run)
		clasificacionGroup.GET("/:id", getByIdController.Run)
		clasificacionGroup.PUT("/:id", updateController.Run)
		clasificacionGroup.DELETE("/:id", deleteController.Run)
		clasificacionGroup.GET("/", getAllController.Run)
		
		// Rutas específicas
		clasificacionGroup.GET("/municipio/:municipio_id", getByMunicipioController.Run)
		
		// Rutas para imágenes
		clasificacionGroup.POST("/:id/upload-image", uploadImageController.Run)
	}

	// Ruta para servir imágenes (fuera del grupo de autenticación si es necesario)
	router.engine.GET("/uploads/images/clasificacion/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		fileUploader.ServeImage(c, "images/clasificacion", filename)
	})
}