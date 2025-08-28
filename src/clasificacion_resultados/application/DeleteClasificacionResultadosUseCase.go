// DeleteClasificacionResultadosUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain/entities"
)

type DeleteClasificacionResultadosUseCase struct {
	repo repositories.IClasificacionResultados
}

func NewDeleteClasificacionResultadosUseCase(repo repositories.IClasificacionResultados) *DeleteClasificacionResultadosUseCase {
	return &DeleteClasificacionResultadosUseCase{repo: repo}
}

func (uc *DeleteClasificacionResultadosUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}

func (uc *DeleteClasificacionResultadosUseCase) GetById(id int32) (*entities.ClasificacionResultados, error) {
	return uc.repo.GetById(id)
}