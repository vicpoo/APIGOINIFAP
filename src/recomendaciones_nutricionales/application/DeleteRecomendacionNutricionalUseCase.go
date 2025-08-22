// En DeleteRecomendacionNutricionalUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type DeleteRecomendacionNutricionalUseCase struct {
	repo repositories.IRecomendacionNutricional
}

func NewDeleteRecomendacionNutricionalUseCase(repo repositories.IRecomendacionNutricional) *DeleteRecomendacionNutricionalUseCase {
	return &DeleteRecomendacionNutricionalUseCase{repo: repo}
}

func (uc *DeleteRecomendacionNutricionalUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}

// Agregar este método nuevo
func (uc *DeleteRecomendacionNutricionalUseCase) GetById(id int32) (*entities.RecomendacionNutricional, error) {
	return uc.repo.GetById(id)
}