// municipio_repository.go
package domain

import (
	"github.com/vicpoo/APIGOINIFAP/src/municipios/domain/entities"
)

type IMunicipio interface {
	Save(municipio *entities.Municipio) error
	Update(municipio *entities.Municipio) error
	Delete(id int32) error
	GetById(id int32) (*entities.Municipio, error)
	GetAll() ([]entities.Municipio, error)
}