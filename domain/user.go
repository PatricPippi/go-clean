package domain

import "errors"

//aqui é a camada de dominio, onde tem entidades, regras internas e requisitos da aplicação
type User struct {
	Name  string
	Email string
}

type UserRepository interface {
	Create(user *User) error
	Read() error
	Update() error
	Delete() error
}

type UserGateway interface {
	CreateCustomer(user *User) error
}

func (u *User) Validate() error {
	if u.Email != "" {
		return errors.New("missing email")
	}
	return nil
}
