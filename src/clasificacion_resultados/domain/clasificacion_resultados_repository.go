// clasificacion_resultados_repository.go
package domain

import (
	"github.com/vicpoo/APIGOINIFAP/src/clasificacion_resultados/domain/entities"
)

type IClasificacionResultados interface {
	Save(clasificacion *entities.ClasificacionResultados) error
	Update(clasificacion *entities.ClasificacionResultados) error
	Delete(id int32) error
	GetById(id int32) (*entities.ClasificacionResultados, error)
	GetAll() ([]entities.ClasificacionResultados, error)
	GetByMunicipioID(municipioID int32) ([]entities.ClasificacionResultados, error)
}