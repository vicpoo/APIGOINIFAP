// GetAllRecomendacionesNutricionalesUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type GetAllRecomendacionesNutricionalesUseCase struct {
	repo repositories.IRecomendacionNutricional
}

func NewGetAllRecomendacionesNutricionalesUseCase(repo repositories.IRecomendacionNutricional) *GetAllRecomendacionesNutricionalesUseCase {
	return &GetAllRecomendacionesNutricionalesUseCase{repo: repo}
}

func (uc *GetAllRecomendacionesNutricionalesUseCase) Run() ([]entities.RecomendacionNutricional, error) {
	return uc.repo.GetAll()
}