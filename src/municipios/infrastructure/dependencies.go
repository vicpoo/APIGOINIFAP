// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/APIGOINIFAP/src/municipios/application"
)

func InitMunicipioDependencies() (
	*CreateMunicipioController,
	*GetMunicipioByIdController,
	*UpdateMunicipioController,
	*DeleteMunicipioController,
	*GetAllMunicipiosController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLMunicipioRepository()

	// Casos de uso
	createUseCase := application.NewCreateMunicipioUseCase(repo)
	getByIdUseCase := application.NewGetMunicipioByIdUseCase(repo)
	updateUseCase := application.NewUpdateMunicipioUseCase(repo)
	deleteUseCase := application.NewDeleteMunicipioUseCase(repo)
	getAllUseCase := application.NewGetAllMunicipiosUseCase(repo)

	// Controladores
	createController := NewCreateMunicipioController(createUseCase)
	getByIdController := NewGetMunicipioByIdController(getByIdUseCase)
	updateController := NewUpdateMunicipioController(updateUseCase)
	deleteController := NewDeleteMunicipioController(deleteUseCase)
	getAllController := NewGetAllMunicipiosController(getAllUseCase)

	return createController, getByIdController, updateController, deleteController, getAllController
}