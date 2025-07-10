package domain_test

import (
	"testing"

	"github.com/i0li/super_shiharai_kun/internal/domain"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestInvoice(t *testing.T) {
	type args struct {
		paymentAmount decimal.Decimal
		feeRate       decimal.Decimal
		taxRate       decimal.Decimal
	}
	type want struct {
		fee         decimal.Decimal
		tax         decimal.Decimal
		totalAmount decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "normal",
			args: args{
				paymentAmount: decimal.NewFromInt(10000),
				feeRate:       decimal.NewFromFloat(0.04),
				taxRate:       decimal.NewFromFloat(0.10),
			},
			want: want{
				fee:         decimal.NewFromFloat(400.00),
				tax:         decimal.NewFromFloat(40.00),
				totalAmount: decimal.NewFromFloat(10440.00),
			},
		},
		{
			name: "complex",
			args: args{
				paymentAmount: decimal.NewFromInt(123456789),
				feeRate:       decimal.NewFromFloat(0.012345),
				taxRate:       decimal.NewFromFloat(0.12345),
			},
			want: want{
				fee:         decimal.NewFromInt(1524074),
				tax:         decimal.NewFromInt(188146),
				totalAmount: decimal.NewFromInt(125169009),
			},
		},
		{
			name: "zero rates",
			args: args{
				paymentAmount: decimal.NewFromInt(500),
				feeRate:       decimal.NewFromFloat(0.00),
				taxRate:       decimal.NewFromFloat(0.00),
			},
			want: want{
				fee:         decimal.NewFromFloat(0.00),
				tax:         decimal.NewFromFloat(0.00),
				totalAmount: decimal.NewFromFloat(500.00),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoice := &domain.Invoice{
				PaymentAmount: tt.args.paymentAmount,
				FeeRate:       tt.args.feeRate,
				TaxRate:       tt.args.taxRate,
			}
			invoice.CalculateTotalAmount()

			require.Equal(t, tt.want.fee.String(), invoice.Fee.String(), "Fee mismatch")
			require.Equal(t, tt.want.tax.String(), invoice.TaxAmount.String(), "Tax mismatch")
			require.Equal(t, tt.want.totalAmount.String(), invoice.TotalAmount.String(), "Total mismatch")
		})
	}
}
