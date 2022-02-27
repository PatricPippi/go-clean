package usecase

import (
	"github.com/go-clean/domain"
)

//camada use case, camada de regras de negocio da aplicação, a camada sempre retorna e recebe um domain
type repository interface {
	Create(user *domain.User) error
	Read() error
	Update() error
	Delete() error
}

type customerService interface {
	CreateCustomer(user *domain.User) error
}

type UserUseCase struct {
	repository      repository
	customerService customerService
}

func NewUserUseCase(userRepository repository, customerService customerService) *UserUseCase {
	return &UserUseCase{
		repository:      userRepository,
		customerService: customerService,
	}
}

func (uc *UserUseCase) CreateUser(user *domain.User) error {
	err := uc.customerService.CreateCustomer(user)
	if err != nil {
		return err
	}
	err = uc.repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}
