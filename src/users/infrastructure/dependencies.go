// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/APIGOINIFAP/src/users/application"
)

func InitUserDependencies() (
	*CreateUserController,
	*GetUserByIdController,
	*UpdateUserController,
	*DeleteUserController,
	*GetAllUsersController,
	*LoginUserController,
) {
	// Repositorio
	repo := NewMySQLUserRepository()

	// Casos de uso
	createUseCase := application.NewCreateUserUseCase(repo)
	getByIdUseCase := application.NewGetUserByIdUseCase(repo)
	updateUseCase := application.NewUpdateUserUseCase(repo)
	deleteUseCase := application.NewDeleteUserUseCase(repo)
	getAllUseCase := application.NewGetAllUsersUseCase(repo)
	loginUseCase := application.NewLoginUserUseCase(repo)

	// Controladores
	createController := NewCreateUserController(createUseCase)
	getByIdController := NewGetUserByIdController(getByIdUseCase)
	updateController := NewUpdateUserController(updateUseCase)
	deleteController := NewDeleteUserController(deleteUseCase)
	getAllController := NewGetAllUsersController(getAllUseCase)
	loginController := NewLoginUserController(loginUseCase)

	return createController, getByIdController, updateController, deleteController, getAllController, loginController
}

