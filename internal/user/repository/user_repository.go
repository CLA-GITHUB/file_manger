package user

import (
	"errors"

	"github.com/cla-github/file_manager/db"
	userModel "github.com/cla-github/file_manager/internal/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserById(userId uuid.UUID) *userModel.User {
	var user userModel.User

	err := db.DB.First(&user, "id = ?", userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}

func GetUserByEmail(email string) *userModel.User {
	var user *userModel.User

	err := db.DB.First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return user
}

func CreateUser(data userModel.User) (*userModel.User, error) {
	var checkUser *userModel.User

	//check if email has been taken.
	checkUser = GetUserByEmail(data.Email)
	if checkUser != nil {
		return nil, errors.New("Email address taken, try logging in.")
	}

	newUser := userModel.User{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
	result := db.DB.Create(&newUser)

	// return any error
	if result.Error != nil {
		return nil, result.Error
	}

	newUser.Password = ""

	return &newUser, nil
}
