//GetUserByIdUseCase.go
package application

import (
	"github.com/vicpoo/APIGOINIFAP/src/users/domain"
	"github.com/vicpoo/APIGOINIFAP/src/users/domain/entities"
)

type GetUserByIdUseCase struct {
	repo domain.IUser
}

func NewGetUserByIdUseCase(repo domain.IUser) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{repo: repo}
}

func (uc *GetUserByIdUseCase) Run(id int32) (*entities.User, error) {
	return uc.repo.GetById(id)
}
