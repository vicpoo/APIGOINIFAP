// GetMunicipioByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/municipios/domain"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/domain/entities"
)

type GetMunicipioByIdUseCase struct {
	repo repositories.IMunicipio
}

func NewGetMunicipioByIdUseCase(repo repositories.IMunicipio) *GetMunicipioByIdUseCase {
	return &GetMunicipioByIdUseCase{repo: repo}
}

func (uc *GetMunicipioByIdUseCase) Run(id int32) (*entities.Municipio, error) {
	return uc.repo.GetById(id)
}
