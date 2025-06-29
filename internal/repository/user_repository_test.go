package repository_test

import (
	"testing"
	"time"

	"github.com/i0li/super_shiharai_kun/internal/domain"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"github.com/stretchr/testify/require"
)

func TestUserRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)

	user := &domain.User{
		CompanyName: "CIN GROUP",
		Name:        "Iori Sakino",
		Email:       "iori@gmail.com",
		Password:    "password",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Create
	err := repo.Create(user)
	require.NoError(t, err)
	require.NotZero(t, user.ID)

	// FindByEmail
	found, err := repo.FindByEmail(user.Email)
	require.NoError(t, err)
	require.Equal(t, user.Email, found.Email)
	require.Equal(t, user.Name, found.Name)
}
