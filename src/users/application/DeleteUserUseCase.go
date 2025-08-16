//DeleteUserUseCase.go
package application

import "github.com/vicpoo/APIGOINIFAP/src/users/domain"

type DeleteUserUseCase struct {
	repo domain.IUser
}

func NewDeleteUserUseCase(repo domain.IUser) *DeleteUserUseCase {
	return &DeleteUserUseCase{repo: repo}
}

func (uc *DeleteUserUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}
