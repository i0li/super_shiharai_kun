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
		PaymentAmount:  decimal.NewFromInt(12345),
		Fee:            decimal.NewFromFloat(123.45),
		FeeRate:        decimal.NewFromFloat(0.01),
		TaxAmount:      decimal.NewFromFloat(18.5175),
		TaxRate:        decimal.NewFromFloat(0.15),
		TotalAmount:    decimal.NewFromFloat(12363.5175),
		PaymentDueDate: time.Date(2025, 7, 5, 0, 0, 0, 0, time.Local),
		CreatedAt:      time.Date(2025, 6, 5, 11, 30, 0, 0, time.Local),
		UpdatedAt:      time.Date(2025, 6, 5, 11, 30, 0, 0, time.Local),
	}
}
