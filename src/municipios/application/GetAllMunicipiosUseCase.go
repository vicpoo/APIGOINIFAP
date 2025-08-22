// GetAllMunicipiosUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/municipios/domain"
	"github.com/vicpoo/APIGOINIFAP/src/municipios/domain/entities"
)

type GetAllMunicipiosUseCase struct {
	repo repositories.IMunicipio
}

func NewGetAllMunicipiosUseCase(repo repositories.IMunicipio) *GetAllMunicipiosUseCase {
	return &GetAllMunicipiosUseCase{repo: repo}
}

func (uc *GetAllMunicipiosUseCase) Run() ([]entities.Municipio, error) {
	return uc.repo.GetAll()
}

