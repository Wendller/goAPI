package entities

import (
	entities "github.com/Wendller/goexpert/apis/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entities.ID `json:"id"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entities.NewID(),
		Name:     name,
		Email:    email,
		Password: string(passwordHash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
