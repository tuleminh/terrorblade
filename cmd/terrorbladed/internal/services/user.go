package services

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"terrorblade"
	"terrorblade/cmd/terrorbladed/internal/dtos"
	_baseDTOs "terrorblade/dtos"
)

func NewUserService(db *gorm.DB, userRepository terrorblade.UserRepository) *UserService {
	return &UserService{
		db:             db,
		userRepository: userRepository,
	}
}

type UserService struct {
	db             *gorm.DB
	userRepository terrorblade.UserRepository
}

func (_this *UserService) CreateUser(request *dtos.CreateUserRequest) (*dtos.CreateUserResponse, error) {
	tx := _this.db.Begin()
	defer tx.RollbackUnlessCommitted()

	encryptedPasswd, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	user := terrorblade.User{
		FullName: request.FullName,
		Username: request.Username,
		Password: string(encryptedPasswd),
		Status:   terrorblade.UserStatus(request.Status),
	}
	err := _this.userRepository.CreateUser(tx, &user)
	if err != nil {
		return nil, err
	}

	return &dtos.CreateUserResponse{
		Metadata: _baseDTOs.Metadata{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: dtos.CreateUserData{
			ID: user.ID,
		},
	}, nil
}

func (_this *UserService) GetUser(id int64) (*dtos.GetUserResponse, error) {
	user, err := _this.userRepository.GetUser(_this.db, id)
	if err != nil {
		return nil, err
	}
	return &dtos.GetUserResponse{
		Metadata: _baseDTOs.Metadata{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: dtos.GetUserData{
			ID:       user.ID,
			FullName: user.FullName,
			Username: user.Username,
			Password: user.Password,
			Status:   string(user.Status),
		},
	}, nil
}
