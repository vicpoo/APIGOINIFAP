// user_repository.go
package domain

import (
	"github.com/vicpoo/APIGOINIFAP/src/users/domain/entities"
)

type IUser interface {
	Save(user *entities.User) error
	Update(user *entities.User) error
	Delete(id int32) error
	GetById(id int32) (*entities.User, error)
	GetAll() ([]entities.User, error)
	Login(correo, password string) (*entities.User, error)
}
