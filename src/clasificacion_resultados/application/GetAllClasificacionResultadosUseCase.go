// GetAllClasificacionResultadosUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain"
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain/entities"
)

type GetAllClasificacionResultadosUseCase struct {
	repo repositories.IClasificacionResultados
}

func NewGetAllClasificacionResultadosUseCase(repo repositories.IClasificacionResultados) *GetAllClasificacionResultadosUseCase {
	return &GetAllClasificacionResultadosUseCase{repo: repo}
}

func (uc *GetAllClasificacionResultadosUseCase) Run() ([]entities.ClasificacionResultados, error) {
	return uc.repo.GetAll()
}