//LoginUserUseCase.go
package application

import (
	"github.com/vicpoo/APIGOINIFAP/src/users/domain"
	"github.com/vicpoo/APIGOINIFAP/src/users/domain/entities"
)

type LoginUserUseCase struct {
	repo domain.IUser
}

func NewLoginUserUseCase(repo domain.IUser) *LoginUserUseCase {
	return &LoginUserUseCase{repo: repo}
}

func (uc *LoginUserUseCase) Run(correo, password string) (*entities.User, error) {
	return uc.repo.Login(correo, password)
}
