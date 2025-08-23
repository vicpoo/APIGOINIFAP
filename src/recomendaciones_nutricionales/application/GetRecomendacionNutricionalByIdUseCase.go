// GetRecomendacionNutricionalByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type GetRecomendacionNutricionalByIdUseCase struct {
	repo repositories.IRecomendacionNutricional
}

func NewGetRecomendacionNutricionalByIdUseCase(repo repositories.IRecomendacionNutricional) *GetRecomendacionNutricionalByIdUseCase {
	return &GetRecomendacionNutricionalByIdUseCase{repo: repo}
}

func (uc *GetRecomendacionNutricionalByIdUseCase) Run(id int32) (*entities.RecomendacionNutricional, error) {
	return uc.repo.GetById(id)
}