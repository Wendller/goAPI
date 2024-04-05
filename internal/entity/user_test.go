package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserCreation(t *testing.T) {
	t.Run("Valid User", func(t *testing.T) {
		user, err := NewUser("John Doe", "john@mail.com", "123456")

		assert.Nil(t, err)
		assert.NotEmpty(t, user.ID)
		assert.NotEmpty(t, user.Password)
		assert.Equal(t, user.Name, "John Doe")
		assert.Equal(t, user.Email, "john@mail.com")
	})

	t.Run("Password Validation", func(t *testing.T) {
		user, err := NewUser("John Doe", "john@mail.com", "123456")

		assert.Nil(t, err)

		t.Run("Valid Password", func(t *testing.T) {
			assert.True(t, user.ValidatePassword("123456"))
		})

		t.Run("Invalid Password", func(t *testing.T) {
			assert.False(t, user.ValidatePassword("1234567"))
		})

		t.Run("Password Encryption", func(t *testing.T) {
			assert.NotEqual(t, user.Password, "123456")
		})
	})
}
