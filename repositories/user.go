package repositories

import (
	"github.com/jinzhu/gorm"

	"terrorblade"
)

// NewUserRepository returns a new instance of UserRepository.
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

type UserRepository struct {
}

func (_this *UserRepository) CreateUser(db *gorm.DB, user *terrorblade.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (_this *UserRepository) GetUser(db *gorm.DB, id int64) (*terrorblade.User, error) {
	var user terrorblade.User

	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
