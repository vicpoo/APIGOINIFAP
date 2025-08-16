package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/rol/domain"
	"github.com/vicpoo/APIGOINIFAP/src/rol/domain/entities"
)

type CreateRolUseCase struct {
	repo repositories.IRol
}

func NewCreateRolUseCase(repo repositories.IRol) *CreateRolUseCase {
	return &CreateRolUseCase{repo: repo}
}

func (uc *CreateRolUseCase) Run(rol *entities.Rol) (*entities.Rol, error) {
	err := uc.repo.Save(rol)
	if err != nil {
		return nil, err
	}
	return rol, nil
}
