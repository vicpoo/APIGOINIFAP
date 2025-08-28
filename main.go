// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOINIFAP/src/core"
	"github.com/vicpoo/APIGOINIFAP/src/rol/infrastructure"
	usersInfrastructure "github.com/vicpoo/APIGOINIFAP/src/users/infrastructure"
	municipiosInfrastructure "github.com/vicpoo/APIGOINIFAP/src/municipios/infrastructure"
	recomendacionesinfrastructure "github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/infrastructure"
	clasificacionInfrastructure "github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/infrastructure"
)

func main() {
	// Inicializar base de datos
	core.InitDB()

	// Inicializar motor de rutas
	engine := gin.Default()

	// Configurar CORS
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Registrar rutas de roles
	rolRouter := infrastructure.NewRolRouter(engine)
	rolRouter.Run()

	// Registrar rutas de usuarios
	userRouter := usersInfrastructure.NewUserRouter(engine)
	userRouter.Run()

	// Registrar rutas de municipios
	municipiosRouter := municipiosInfrastructure.NewMunicipioRouter(engine)
	municipiosRouter.Run()

	// Registrar rutas de recomendaciones
	recomendacionesRouter := recomendacionesinfrastructure.NewRecomendacionNutricionalRouter(engine)
	recomendacionesRouter.Run()

	// Registrar rutas de clasificacion
	clasificacionRouter := clasificacionInfrastructure.NewClasificacionResultadosRouter(engine)
	clasificacionRouter.Run()

	// Correr servidor en el puerto 8000
	port := ":8000"
	fmt.Println("Servidor corriendo en http://localhost" + port)
	if err := engine.Run(port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
