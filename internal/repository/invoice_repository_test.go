package repository_test

import (
	"testing"
	"time"

	"github.com/i0li/super_shiharai_kun/internal/domain"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestInvoiceRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewInvoiceRepository(db)

	user := &domain.User{
		CompanyName: "test corp",
		Name:        "test user",
		Email:       "test@example.com",
		Password:    "password",
	}
	err := db.Create(user).Error
	require.NoError(t, err)

	invoice1 := &domain.Invoice{
		UserID:         1,
		IssueDate:      time.Now(),
		PaymentAmount:  decimal.NewFromInt(10000),
		Fee:            decimal.NewFromInt(400),
		FeeRate:        decimal.NewFromFloat(0.04),
		TaxAmount:      decimal.NewFromInt(40),
		TaxRate:        decimal.NewFromFloat(0.1),
		TotalAmount:    decimal.NewFromInt(10440),
		PaymentDueDate: atMidnight(time.Now().AddDate(0, 0, 10)),
	}
	invoice2 := &domain.Invoice{
		UserID:         1,
		IssueDate:      time.Now(),
		PaymentAmount:  decimal.NewFromInt(20000),
		Fee:            decimal.NewFromInt(800),
		FeeRate:        decimal.NewFromFloat(0.04),
		TaxAmount:      decimal.NewFromInt(80),
		TaxRate:        decimal.NewFromFloat(0.1),
		TotalAmount:    decimal.NewFromInt(20880),
		PaymentDueDate: atMidnight(time.Now().AddDate(0, 0, 15)),
	}

	// Create
	err = repo.Create(invoice1)
	require.NoError(t, err)
	err = repo.Create(invoice2)
	require.NoError(t, err)

	// FindByPaymentDueDatePeriod
	start := atMidnight(time.Now().AddDate(0, 0, 10))
	end := atMidnight(time.Now().AddDate(0, 0, 15))
	invoices, err := repo.FindByPaymentDueDatePeriod(1, start, end, 10, 0)
	require.NoError(t, err)
	require.Len(t, invoices, 2)
	require.True(t, invoices[0].PaymentDueDate.After(invoices[1].PaymentDueDate))

	// CountByPaymentDueDatePeriod
	total, err := repo.CountByPaymentDueDatePeriod(1, start, end)
	require.NoError(t, err)
	require.Equal(t, int64(2), total)
}

func atMidnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
