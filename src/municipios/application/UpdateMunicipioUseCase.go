// UpdateMunicipioUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/municipios/domain"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/domain/entities"
)

type UpdateMunicipioUseCase struct {
	repo repositories.IMunicipio
}

func NewUpdateMunicipioUseCase(repo repositories.IMunicipio) *UpdateMunicipioUseCase {
	return &UpdateMunicipioUseCase{repo: repo}
}

func (uc *UpdateMunicipioUseCase) Run(municipio *entities.Municipio) (*entities.Municipio, error) {
	err := uc.repo.Update(municipio)
	if err != nil {
		return nil, err
	}
	return municipio, nil
}