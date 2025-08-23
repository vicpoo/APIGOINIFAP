// DownloadRecomendacionNutricionalUseCase.go
package application

import (
	repositories "github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain"
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type DownloadRecomendacionNutricionalUseCase struct {
	repo repositories.IRecomendacionNutricional
}

func NewDownloadRecomendacionNutricionalUseCase(repo repositories.IRecomendacionNutricional) *DownloadRecomendacionNutricionalUseCase {
	return &DownloadRecomendacionNutricionalUseCase{repo: repo}
}

func (uc *DownloadRecomendacionNutricionalUseCase) RunByID(id int32) (*entities.RecomendacionNutricional, error) {
	return uc.repo.GetById(id)
}

func (uc *DownloadRecomendacionNutricionalUseCase) RunByMunicipio(municipioID int32) ([]entities.RecomendacionNutricional, error) {
	return uc.repo.GetByMunicipioID(municipioID)
}