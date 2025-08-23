// recomendacion_nutricional_repository.go
package domain

import (
	"github.com/vicpoo/APIGOINIFAP/src/recomendaciones_nutricionales/domain/entities"
)

type IRecomendacionNutricional interface {
	Save(recomendacion *entities.RecomendacionNutricional) error
	Update(recomendacion *entities.RecomendacionNutricional) error
	Delete(id int32) error
	GetById(id int32) (*entities.RecomendacionNutricional, error)
	GetAll() ([]entities.RecomendacionNutricional, error)
	GetByMunicipioID(municipioID int32) ([]entities.RecomendacionNutricional, error) // Nuevo método
}