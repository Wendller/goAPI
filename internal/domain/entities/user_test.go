package entities

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

}
