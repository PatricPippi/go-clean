package usecase

import (
	"github.com/go-clean/domain"
)

//camada use case, camada de regras de negocio da aplicação, a camada sempre retorna e recebe um domain
type UserUseCase struct {
	repository domain.UserRepository
	gateway    domain.UserGateway
}

func NewUserUseCase(userRepository domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: userRepository,
	}
}

func (u *UserUseCase) CreateUser(user *domain.User) error {
	err := u.gateway.CreateCustomer(user)
	if err != nil {
		return err
	}
	err = u.repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

//TODO: Rest of user functions...
