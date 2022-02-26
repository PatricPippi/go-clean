package domain

import "errors"

//aqui é a camada de dominio, onde tem entidades e validações internas da aplicação

type User struct {
	name  string
	email string
}

func (u *User) Validate() error {
	if u.email != "" {
		return errors.New("missing email")
	}
	return nil
}
