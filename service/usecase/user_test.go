package usecase

import (
	"testing"

	"github.com/go-clean/service/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type userRepoMock struct {
	mock.Mock
}

func (u *userRepoMock) Create(user *domain.User) error {
	return nil
}
func (u *userRepoMock) Read() error {
	return nil
}
func (u *userRepoMock) Update() error {
	return nil
}
func (u *userRepoMock) Delete() error {
	return nil
}
func TestCreateUser(t *testing.T) {
	userRepoMock := new(userRepoMock)
	userRepoMock.On("Create", mock.AnythingOfType("*domain.User")).Return(nil)
	userUseCase := NewUserUseCase(userRepoMock)
	err := userUseCase.CreateUser(&domain.User{
		Name:  "test",
		Email: "test",
	})
	assert.NoError(t, err)
}
