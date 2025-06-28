package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthUsecase interface {
	Login(email, password string) (int64, error)
	GenerateToken(userID int64) (string, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{userRepo: userRepo}
}

func (uc *authUsecase) Login(email, password string) (int64, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, apperr.Wrap(apperr.ErrUnauthorized)
	} else if err != nil {
		return 0, apperr.Wrap(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return 0, apperr.Wrap(apperr.ErrUnauthorized)
	}

	return user.ID, nil
}

func (uc *authUsecase) GenerateToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": fmt.Sprintf("%d", userID),
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", apperr.Wrap(err)
	}
	return tokenString, nil
}
