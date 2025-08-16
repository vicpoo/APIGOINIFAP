//GetAllUsersUseCase.go
package application

import (
	"github.com/vicpoo/APIGOINIFAP/src/users/domain"
	"github.com/vicpoo/APIGOINIFAP/src/users/domain/entities"
)

type GetAllUsersUseCase struct {
	repo domain.IUser
}

func NewGetAllUsersUseCase(repo domain.IUser) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{repo: repo}
}

func (uc *GetAllUsersUseCase) Run() ([]entities.User, error) {
	return uc.repo.GetAll()
}
