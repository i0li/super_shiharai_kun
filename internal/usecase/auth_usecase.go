package usecase

import (
	"github.com/cockroachdb/errors"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	jwtUtil "github.com/i0li/super_shiharai_kun/internal/util"
)

type AuthUsecase interface {
	Login(email, password string) (string, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{userRepo: userRepo}
}

func (uc *authUsecase) Login(email, password string) (string, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", apperr.Wrap(apperr.ErrUnauthorized)
	} else if err != nil {
		return "", apperr.Wrap(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", apperr.Wrap(apperr.ErrUnauthorized)
	}

	tokenStr, err := jwtUtil.GenerateToken(user.ID)
	if err != nil {
		return "", apperr.Wrap(err)
	}

	return tokenStr, nil
}
