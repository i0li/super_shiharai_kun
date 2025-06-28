package usecase

import (
	"github.com/cockroachdb/errors"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/model"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	_, err := uc.repo.FindByEmail(user.Email)
	if err == nil {
		return apperr.Wrap(apperr.ErrEmailAlreadyExists)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return apperr.Wrap(err)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperr.Wrap(err)
	}
	user.Password = string(hashed)

	return uc.repo.Create(user)
}
