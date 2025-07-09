package testdata

import (
	"time"

	"github.com/i0li/super_shiharai_kun/internal/domain"
	"github.com/shopspring/decimal"
)

func StandardInvoice() *domain.Invoice {
	return &domain.Invoice{
		ID:             1,
		UserID:         1,
		IssueDate:      time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local),
		PaymentAmount:  decimal.NewFromInt(10000),
		Fee:            decimal.NewFromFloat(400),
		FeeRate:        decimal.NewFromFloat(0.04),
		TaxAmount:      decimal.NewFromFloat(40),
		TaxRate:        decimal.NewFromFloat(0.10),
		TotalAmount:    decimal.NewFromFloat(10440),
		PaymentDueDate: time.Date(2025, 7, 1, 0, 0, 0, 0, time.Local),
		CreatedAt:      time.Date(2025, 6, 1, 10, 0, 0, 0, time.Local),
		UpdatedAt:      time.Date(2025, 6, 1, 10, 0, 0, 0, time.Local),
	}
}

func ComplexCalculatedInvoice() *domain.Invoice {
	return &domain.Invoice{
		ID:             2,
		UserID:         1,
		IssueDate:      time.Date(2025, 6, 5, 0, 0, 0, 0, time.Local),
		PaymentAmount:  decimal.NewFromInt(123456789),
		Fee:            decimal.NewFromInt(1524074),
		FeeRate:        decimal.NewFromFloat(0.012345),
		TaxAmount:      decimal.NewFromFloat(188146),
		TaxRate:        decimal.NewFromFloat(0.12345),
		TotalAmount:    decimal.NewFromFloat(125169009),
		PaymentDueDate: time.Date(2025, 7, 5, 0, 0, 0, 0, time.Local),
		CreatedAt:      time.Date(2025, 6, 5, 11, 30, 0, 0, time.Local),
		UpdatedAt:      time.Date(2025, 6, 5, 11, 30, 0, 0, time.Local),
	}
}
