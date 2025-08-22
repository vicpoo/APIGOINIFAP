// UpdateRecomendacionNutricionalUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type UpdateRecomendacionNutricionalUseCase struct {
	repo repositories.IRecomendacionNutricional
}

func NewUpdateRecomendacionNutricionalUseCase(repo repositories.IRecomendacionNutricional) *UpdateRecomendacionNutricionalUseCase {
	return &UpdateRecomendacionNutricionalUseCase{repo: repo}
}

func (uc *UpdateRecomendacionNutricionalUseCase) Run(recomendacion *entities.RecomendacionNutricional) (*entities.RecomendacionNutricional, error) {
	err := uc.repo.Update(recomendacion)
	if err != nil {
		return nil, err
	}
	return recomendacion, nil
}

// Agregar este método nuevo para obtener por ID
func (uc *UpdateRecomendacionNutricionalUseCase) GetById(id int32) (*entities.RecomendacionNutricional, error) {
	return uc.repo.GetById(id)
}