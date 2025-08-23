// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/application"
)

func InitRecomendacionNutricionalDependencies() (
	*CreateRecomendacionNutricionalController,
	*GetRecomendacionNutricionalByIdController,
	*UpdateRecomendacionNutricionalController,
	*DeleteRecomendacionNutricionalController,
	*GetAllRecomendacionesNutricionalesController,
	*DownloadRecomendacionNutricionalController, // Nuevo
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLRecomendacionNutricionalRepository()

	// Casos de uso
	createUseCase := application.NewCreateRecomendacionNutricionalUseCase(repo)
	getByIdUseCase := application.NewGetRecomendacionNutricionalByIdUseCase(repo)
	updateUseCase := application.NewUpdateRecomendacionNutricionalUseCase(repo)
	deleteUseCase := application.NewDeleteRecomendacionNutricionalUseCase(repo)
	getAllUseCase := application.NewGetAllRecomendacionesNutricionalesUseCase(repo)
	downloadUseCase := application.NewDownloadRecomendacionNutricionalUseCase(repo) // Nuevo

	// Controladores
	createController := NewCreateRecomendacionNutricionalController(createUseCase)
	getByIdController := NewGetRecomendacionNutricionalByIdController(getByIdUseCase)
	updateController := NewUpdateRecomendacionNutricionalController(updateUseCase)
	deleteController := NewDeleteRecomendacionNutricionalController(deleteUseCase)
	getAllController := NewGetAllRecomendacionesNutricionalesController(getAllUseCase)
	downloadController := NewDownloadRecomendacionNutricionalController(downloadUseCase) // Nuevo

	return createController, getByIdController, updateController, deleteController, getAllController, downloadController
}