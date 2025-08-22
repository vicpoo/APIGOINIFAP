// DeleteMunicipioUseCase.go
package application

import repositories "github.com/vicpoo/APIGOINIFAP/src/municipios/domain"

type DeleteMunicipioUseCase struct {
	repo repositories.IMunicipio
}

func NewDeleteMunicipioUseCase(repo repositories.IMunicipio) *DeleteMunicipioUseCase {
	return &DeleteMunicipioUseCase{repo: repo}
}

func (uc *DeleteMunicipioUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}