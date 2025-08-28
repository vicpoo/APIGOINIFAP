// UpdateClasificacionResultadosUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain/entities"
)

type UpdateClasificacionResultadosUseCase struct {
	repo repositories.IClasificacionResultados
}

func NewUpdateClasificacionResultadosUseCase(repo repositories.IClasificacionResultados) *UpdateClasificacionResultadosUseCase {
	return &UpdateClasificacionResultadosUseCase{repo: repo}
}

func (uc *UpdateClasificacionResultadosUseCase) Run(clasificacion *entities.ClasificacionResultados) (*entities.ClasificacionResultados, error) {
	err := uc.repo.Update(clasificacion)
	if err != nil {
		return nil, err
	}
	return clasificacion, nil
}

func (uc *UpdateClasificacionResultadosUseCase) GetById(id int32) (*entities.ClasificacionResultados, error) {
	return uc.repo.GetById(id)
}