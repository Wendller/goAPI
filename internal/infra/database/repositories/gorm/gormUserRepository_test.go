package gorm_repositories

import (
	"testing"

	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestInsertUser(t *testing.T) {
	t.Run("Insert User", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		db.AutoMigrate(&entities.User{})
		user, _ := entities.NewUser("John Doe", "john@mail.com", "123456")
		userRepository := NewUserRepository(db)

		err = userRepository.Create(user)
		assert.Nil(t, err)

		var createdUser entities.User
		err = db.First(&createdUser, "id = ?", user.ID).Error

		assert.Nil(t, err)
		assert.Equal(t, createdUser.ID, user.ID)
	})
}

func TestFindUserByEmail(t *testing.T) {
	t.Run("Return user by email", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		db.AutoMigrate(&entities.User{})
		user, _ := entities.NewUser("John Doe", "john@mail.com", "123456")
		userRepository := NewUserRepository(db)

		err = userRepository.Create(user)
		assert.Nil(t, err)

		createdUser, err := userRepository.FindByEmail(user.Email)

		assert.Nil(t, err)
		assert.Equal(t, createdUser.ID, user.ID)
	})
}
