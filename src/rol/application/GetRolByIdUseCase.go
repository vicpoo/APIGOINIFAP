//GetRolByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/rol/domain"
	"github.com/vicpoo/APIGOINIFAP/src/rol/domain/entities"
)

type GetRolByIdUseCase struct {
	repo repositories.IRol
}

func NewGetRolByIdUseCase(repo repositories.IRol) *GetRolByIdUseCase {
	return &GetRolByIdUseCase{repo: repo}
}

func (uc *GetRolByIdUseCase) Run(id int32) (*entities.Rol, error) {
	return uc.repo.GetById(id)
}
