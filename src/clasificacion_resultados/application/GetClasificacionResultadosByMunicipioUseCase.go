// GetClasificacionResultadosByMunicipioUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain/entities"
)

type GetClasificacionResultadosByMunicipioUseCase struct {
	repo repositories.IClasificacionResultados
}

func NewGetClasificacionResultadosByMunicipioUseCase(repo repositories.IClasificacionResultados) *GetClasificacionResultadosByMunicipioUseCase {
	return &GetClasificacionResultadosByMunicipioUseCase{repo: repo}
}

func (uc *GetClasificacionResultadosByMunicipioUseCase) Run(municipioID int32) ([]entities.ClasificacionResultados, error) {
	return uc.repo.GetByMunicipioID(municipioID)
}