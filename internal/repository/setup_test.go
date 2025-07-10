package repository_test

import (
	"testing"

	"github.com/i0li/super_shiharai_kun/internal/domain"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var models = []interface{}{
	&domain.Invoice{},
	&domain.User{},
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.Migrator().DropTable(models...)
	require.NoError(t, err)
	err = db.AutoMigrate(models...)
	require.NoError(t, err)

	return db
}
