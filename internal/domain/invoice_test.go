package domain_test

import (
	"testing"

	"github.com/i0li/super_shiharai_kun/internal/domain"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestInvoice(t *testing.T) {
	tests := []struct {
		name          string
		paymentAmount decimal.Decimal
		feeRate       decimal.Decimal
		taxRate       decimal.Decimal
		expectedFee   decimal.Decimal
		expectedTax   decimal.Decimal
		expectedTotal decimal.Decimal
	}{
		{
			name:          "normal1",
			paymentAmount: decimal.NewFromInt(10000),
			feeRate:       decimal.NewFromFloat(0.04),
			taxRate:       decimal.NewFromFloat(0.10),
			expectedFee:   decimal.NewFromFloat(400.00),
			expectedTax:   decimal.NewFromFloat(40.00),
			expectedTotal: decimal.NewFromFloat(10440.00),
		},
		{
			name:          "normal2",
			paymentAmount: decimal.NewFromInt(1234),
			feeRate:       decimal.NewFromFloat(0.12),
			taxRate:       decimal.NewFromFloat(0.08),
			expectedFee:   decimal.RequireFromString("148.08"),    // 1234 × 0.12
			expectedTax:   decimal.RequireFromString("11.8464"),   // 148.08 × 0.08
			expectedTotal: decimal.RequireFromString("1393.9264"), // 1234 + 148.08 + 11.8464
		},
		{
			name:          "zero rates",
			paymentAmount: decimal.NewFromInt(500),
			feeRate:       decimal.NewFromFloat(0.00),
			taxRate:       decimal.NewFromFloat(0.00),
			expectedFee:   decimal.NewFromFloat(0.00),
			expectedTax:   decimal.NewFromFloat(0.00),
			expectedTotal: decimal.NewFromFloat(500.00),
		},
	}

	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			invoice := &domain.Invoice{
				PaymentAmount: tCase.paymentAmount,
				FeeRate:       tCase.feeRate,
				TaxRate:       tCase.taxRate,
			}
			invoice.CalculateTotalAmount()

			require.Equal(t, tCase.expectedFee.String(), invoice.Fee.String(), "Fee mismatch")
			require.Equal(t, tCase.expectedTax.String(), invoice.TaxAmount.String(), "Tax mismatch")
			require.Equal(t, tCase.expectedTotal.String(), invoice.TotalAmount.String(), "Total mismatch")
		})
	}
}
