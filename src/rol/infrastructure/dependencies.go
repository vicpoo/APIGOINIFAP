package infrastructure

import (
	"github.com/vicpoo/APIGOINIFAP/src/rol/application"
)

func InitRolDependencies() (
	*CreateRolController,
	*GetRolByIdController,
	*UpdateRolController,
	*DeleteRolController,
	*GetAllRolesController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLRolRepository()

	// Casos de uso
	createUseCase := application.NewCreateRolUseCase(repo)
	getByIdUseCase := application.NewGetRolByIdUseCase(repo)
	updateUseCase := application.NewUpdateRolUseCase(repo)
	deleteUseCase := application.NewDeleteRolUseCase(repo)
	getAllUseCase := application.NewGetAllRolesUseCase(repo)

	// Controladores
	createController := NewCreateRolController(createUseCase)
	getByIdController := NewGetRolByIdController(getByIdUseCase)
	updateController := NewUpdateRolController(updateUseCase)
	deleteController := NewDeleteRolController(deleteUseCase)
	getAllController := NewGetAllRolesController(getAllUseCase)

	return createController, getByIdController, updateController, deleteController, getAllController
}
