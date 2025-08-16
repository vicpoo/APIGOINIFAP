package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/rol/domain"
	"github.com/vicpoo/APIGOINIFAP/src/rol/domain/entities"
)

type GetAllRolesUseCase struct {
	repo repositories.IRol
}

func NewGetAllRolesUseCase(repo repositories.IRol) *GetAllRolesUseCase {
	return &GetAllRolesUseCase{repo: repo}
}

func (uc *GetAllRolesUseCase) Run() ([]entities.Rol, error) {
	return uc.repo.GetAll()
}
