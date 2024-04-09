package entities

import (
	"github.com/Wendller/goexpert/goAPI/pkg/entity"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
