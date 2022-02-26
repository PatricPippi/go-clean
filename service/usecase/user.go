package usecase

import (
	"github.com/go-clean/service/domain"
)

//camada use case, camada de regras de negocio da aplicação, a camada sempre retorna e recebe um domain

type UserRepository interface {
	Create(user *domain.User) error
	Read() error
	Update() error
	Delete() error
}

type UserUseCase struct {
	repository UserRepository
}

func NewUserUseCase(userRepository UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: userRepository,
	}
}

func (u *UserUseCase) CreateUser(user *domain.User) error {
	err := u.repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

//TODO: Rest of user functions...
