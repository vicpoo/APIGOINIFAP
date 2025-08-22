// CreateMunicipioUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/municipios/domain"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/domain/entities"
)

type CreateMunicipioUseCase struct {
	repo repositories.IMunicipio
}

func NewCreateMunicipioUseCase(repo repositories.IMunicipio) *CreateMunicipioUseCase {
	return &CreateMunicipioUseCase{repo: repo}
}

func (uc *CreateMunicipioUseCase) Run(municipio *entities.Municipio) (*entities.Municipio, error) {
	err := uc.repo.Save(municipio)
	if err != nil {
		return nil, err
	}
	return municipio, nil
}