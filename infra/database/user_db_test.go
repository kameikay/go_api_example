package database

import (
	"testing"

	"github.com/kameikay/api_example/internal/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}

	db.AutoMigrate(&entities.User{})
	user, err := entities.NewUser("John Doe", "john.doe@email.com", "123456")
	assert.Nil(t, err)
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entities.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound, user.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}

	db.AutoMigrate(&entities.User{})

	user, err := entities.NewUser("John Doe", "john.doe@email.com", "123456")
	assert.Nil(t, err)

	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound, user.Password)
}
