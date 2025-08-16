//CreateUserUseCase.go
package application

import (
	"github.com/vicpoo/APIGOINIFAP/src/users/domain"
	"github.com/vicpoo/APIGOINIFAP/src/users/domain/entities"
)

type CreateUserUseCase struct {
	repo domain.IUser
}

func NewCreateUserUseCase(repo domain.IUser) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Run(user *entities.User) (*entities.User, error) {
	err := uc.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
