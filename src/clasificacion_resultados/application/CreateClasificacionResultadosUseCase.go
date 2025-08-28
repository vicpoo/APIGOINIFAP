// CreateClasificacionResultadosUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain/entities"
)

type CreateClasificacionResultadosUseCase struct {
	repo repositories.IClasificacionResultados
}

func NewCreateClasificacionResultadosUseCase(repo repositories.IClasificacionResultados) *CreateClasificacionResultadosUseCase {
	return &CreateClasificacionResultadosUseCase{repo: repo}
}

func (uc *CreateClasificacionResultadosUseCase) Run(clasificacion *entities.ClasificacionResultados) (*entities.ClasificacionResultados, error) {
	err := uc.repo.Save(clasificacion)
	if err != nil {
		return nil, err
	}
	return clasificacion, nil
}