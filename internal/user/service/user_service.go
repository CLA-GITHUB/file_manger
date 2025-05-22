package user

import (
	model "github.com/cla-github/file_manager/internal/user"
	userRepo "github.com/cla-github/file_manager/internal/user/repository"
	"github.com/google/uuid"
)

func CreateUser(data model.User) error {
	_, err := userRepo.CreateUser(data)

	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) *model.User {
	foundUser := userRepo.GetUserByEmail(email)

	return foundUser
}

func GetMe(uid uuid.UUID) *model.User {
	user := userRepo.GetUserById(uid)

	return user
}
