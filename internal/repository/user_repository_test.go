package repository_test

import (
	"testing"

	"github.com/i0li/super_shiharai_kun/internal/repository"
	"github.com/i0li/super_shiharai_kun/internal/testdata"
	"github.com/stretchr/testify/require"
)

func TestUserRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)

	user := testdata.Alice()

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
