package jwt

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
)

var (
	secretKey       = os.Getenv("JWT_SECRET")
	tokenExpiration = time.Minute * 30
)

func GenerateToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": fmt.Sprintf("%d", userID),
		"exp": time.Now().Add(tokenExpiration).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", apperr.Wrap(err)
	}
	return tokenStr, nil
}

func VerifyToken(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperr.Wrap(errors.New("unexpected signing method"))
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, apperr.Wrap(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, apperr.Wrap(errors.New("invalid token"))
	}

	expVal, ok := claims["exp"]
	if !ok {
		return 0, apperr.Wrap(errors.New("failed to convert exp"))
	}

	expFloat, ok := expVal.(float64)
	if !ok {
		return 0, apperr.Wrap(errors.New("failed to convert exp"))
	}

	expTime := time.Unix(int64(expFloat), 0)
	now := time.Now()
	if now.After(expTime) {
		return 0, apperr.Wrap(errors.New("failed to convert exp"))
	}

	subVal, ok := claims["sub"]
	if !ok {
		return 0, apperr.Wrap(errors.New("failed to convert sub"))
	}

	subStr, ok := subVal.(string)
	if !ok {
		return 0, apperr.Wrap(errors.New("failed to convert sub"))
	}

	userID, err := strconv.ParseInt(subStr, 10, 64)
	if err != nil {
		return 0, apperr.Wrap(errors.New("failed to convert userID"))
	}

	if time.Now().Unix() > int64(claims["exp"].(float64)) {
		return 0, apperr.Wrap(errors.New("token expired"))
	}

	return userID, nil
}
