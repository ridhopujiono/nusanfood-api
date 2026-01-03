package user

import (
	"errors"

	"github.com/ridhopujiono/nusanfood-api/internal/database"
	"gorm.io/gorm"
)

func FindByEmail(email string) (*User, error) {
	var user User

	err := database.DB.
		Where("email = ?", email).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return &user, err
}

func Create(user *User) error {
	return database.DB.Create(user).Error
}
