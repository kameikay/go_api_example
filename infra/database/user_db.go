package database

import (
	"github.com/kameikay/api_example/internal/entities"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) Create(user *entities.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	if err := u.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
