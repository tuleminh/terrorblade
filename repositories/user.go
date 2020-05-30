package repositories

import (
	"github.com/jinzhu/gorm"

	"terrorblade"
)

// NewUserRepository returns a new instance of UserRepository.
func NewUserRepository() terrorblade.UserRepository {
	return &userRepository{}
}

type userRepository struct {
}

func (_this *userRepository) CreateUser(db *gorm.DB, user *terrorblade.User) error {
	return nil
}

func (_this *userRepository) GetUser(db *gorm.DB, id int64) (*terrorblade.User, error) {
	var user terrorblade.User

	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
