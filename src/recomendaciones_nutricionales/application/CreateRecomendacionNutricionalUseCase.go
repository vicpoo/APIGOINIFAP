// CreateRecomendacionNutricionalUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type CreateRecomendacionNutricionalUseCase struct {
	repo repositories.IRecomendacionNutricional
}

func NewCreateRecomendacionNutricionalUseCase(repo repositories.IRecomendacionNutricional) *CreateRecomendacionNutricionalUseCase {
	return &CreateRecomendacionNutricionalUseCase{repo: repo}
}

func (uc *CreateRecomendacionNutricionalUseCase) Run(recomendacion *entities.RecomendacionNutricional) (*entities.RecomendacionNutricional, error) {
	err := uc.repo.Save(recomendacion)
	if err != nil {
		return nil, err
	}
	return recomendacion, nil
}

