package usecase

import (
	"github.com/cockroachdb/errors"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/domain"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase interface {
	RegisterUser(companyName, name, email, password string) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (uc *userUsecase) RegisterUser(companyName, name, email, password string) error {
	_, err := uc.repo.FindByEmail(email)
	if err == nil {
		return apperr.Wrap(apperr.ErrEmailAlreadyExists)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return apperr.Wrap(err)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return apperr.Wrap(err)
	}
	password = string(hashed)

	user := &domain.User{
		CompanyName: companyName,
		Name:        name,
		Email:       email,
		Password:    password,
	}

	return uc.repo.Create(user)
}
