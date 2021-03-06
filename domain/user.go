package domain

import "errors"

//aqui é a camada de dominio, onde tem entidades, regras internas e requisitos da aplicação
type User struct {
	Name  string
	Email string
}

func (u *User) Validate() error {
	if u.Email != "" {
		return errors.New("missing email")
	}
	return nil
}
