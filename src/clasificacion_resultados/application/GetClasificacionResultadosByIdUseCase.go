// GetClasificacionResultadosByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain/entities"
)

type GetClasificacionResultadosByIdUseCase struct {
	repo repositories.IClasificacionResultados
}

func NewGetClasificacionResultadosByIdUseCase(repo repositories.IClasificacionResultados) *GetClasificacionResultadosByIdUseCase {
	return &GetClasificacionResultadosByIdUseCase{repo: repo}
}

func (uc *GetClasificacionResultadosByIdUseCase) Run(id int32) (*entities.ClasificacionResultados, error) {
	return uc.repo.GetById(id)
}