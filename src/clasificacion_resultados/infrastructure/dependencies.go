// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/application"
)

func InitClasificacionResultadosDependencies() (
	*CreateClasificacionResultadosController,
	*GetClasificacionResultadosByIdController,
	*UpdateClasificacionResultadosController,
	*DeleteClasificacionResultadosController,
	*GetAllClasificacionResultadosController,
	*GetClasificacionResultadosByMunicipioController,
	*UploadImageClasificacionResultadosController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLClasificacionResultadosRepository()

	// Casos de uso
	createUseCase := application.NewCreateClasificacionResultadosUseCase(repo)
	getByIdUseCase := application.NewGetClasificacionResultadosByIdUseCase(repo)
	updateUseCase := application.NewUpdateClasificacionResultadosUseCase(repo)
	deleteUseCase := application.NewDeleteClasificacionResultadosUseCase(repo)
	getAllUseCase := application.NewGetAllClasificacionResultadosUseCase(repo)
	getByMunicipioUseCase := application.NewGetClasificacionResultadosByMunicipioUseCase(repo)
	uploadImageUseCase := application.NewUploadImageClasificacionResultadosUseCase(repo)

	// Controladores
	createController := NewCreateClasificacionResultadosController(createUseCase)
	getByIdController := NewGetClasificacionResultadosByIdController(getByIdUseCase)
	updateController := NewUpdateClasificacionResultadosController(updateUseCase)
	deleteController := NewDeleteClasificacionResultadosController(deleteUseCase)
	getAllController := NewGetAllClasificacionResultadosController(getAllUseCase)
	getByMunicipioController := NewGetClasificacionResultadosByMunicipioController(getByMunicipioUseCase)
	uploadImageController := NewUploadImageClasificacionResultadosController(uploadImageUseCase)

	return createController, getByIdController, updateController, deleteController, getAllController, getByMunicipioController, uploadImageController
}