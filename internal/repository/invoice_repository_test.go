package repository_test

import (
	"testing"
	"time"

	"github.com/i0li/super_shiharai_kun/internal/repository"
	"github.com/i0li/super_shiharai_kun/testdata"
	"github.com/stretchr/testify/require"
)

func TestInvoiceRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewInvoiceRepository(db)

	user := testdata.Alice()
	err := db.Create(user).Error
	require.NoError(t, err)

	invoice1 := testdata.StandardInvoice()
	invoice2 := testdata.ComplexCalculatedInvoice()

	// Create
	err = repo.Create(invoice1)
	require.NoError(t, err)
	err = repo.Create(invoice2)
	require.NoError(t, err)

	// FindByPaymentDueDatePeriod
	start := time.Date(2025, 7, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2025, 7, 5, 0, 0, 0, 0, time.Local)
	invoices, err := repo.FindByPaymentDueDatePeriod(1, start, end, 10, 0)
	require.NoError(t, err)
	require.Len(t, invoices, 2)
	require.True(t, invoices[0].PaymentDueDate.After(invoices[1].PaymentDueDate))

	// CountByPaymentDueDatePeriod
	total, err := repo.CountByPaymentDueDatePeriod(1, start, end)
	require.NoError(t, err)
	require.Equal(t, int64(2), total)
}
