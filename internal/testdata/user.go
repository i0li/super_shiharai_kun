package testdata

import (
	"time"

	"github.com/i0li/super_shiharai_kun/internal/domain"
)

func Alice() *domain.User {
	return &domain.User{
		CompanyName: "alice corp",
		Name:        "alice",
		Email:       "alice@gmail.com",
		Password:    "iamalice",
		CreatedAt:   time.Date(2025, 7, 5, 9, 12, 20, 49, time.Local),
		UpdatedAt:   time.Date(2025, 7, 5, 9, 12, 20, 49, time.Local),
	}
}

func Tom() *domain.User {
	return &domain.User{
		CompanyName: "tom corp",
		Name:        "tom",
		Email:       "tom@gmail.com",
		Password:    "iamtom",
		CreatedAt:   time.Date(2025, 7, 5, 9, 18, 30, 34, time.Local),
		UpdatedAt:   time.Date(2025, 7, 5, 9, 18, 30, 34, time.Local),
	}
}
