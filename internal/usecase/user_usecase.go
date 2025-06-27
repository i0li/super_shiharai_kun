package usecase

import (
	appErr "github.com/i0li/super_shiharai_kun/internal/errors"
	"github.com/i0li/super_shiharai_kun/internal/model"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	RegisterUser(user *model.User) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (uc *userUsecase) RegisterUser(user *model.User) error {
	if _, err := uc.repo.FindByEmail(user.Email); err == nil {
		return appErr.ErrEmailAlreadyExists
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	return uc.repo.Create(user)
}
