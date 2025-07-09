package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Invoice struct {
	ID             int64           `gorm:"primaryKey" json:"id"`
	UserID         int64           `gorm:"not null" json:"user_id"`
	IssueDate      time.Time       `gorm:"not null" json:"issue_date"`
	PaymentAmount  decimal.Decimal `gorm:"not null" json:"payment_amount"`
	Fee            decimal.Decimal `gorm:"not null" json:"fee"`
	FeeRate        decimal.Decimal `gorm:"not null" json:"fee_rate"`
	TaxAmount      decimal.Decimal `gorm:"not null" json:"tax_amount"`
	TaxRate        decimal.Decimal `gorm:"not null" json:"tax_rate"`
	TotalAmount    decimal.Decimal `gorm:"not null" json:"total_amount"`
	PaymentDueDate time.Time       `gorm:"not null" json:"payment_due_date"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

func (i *Invoice) CalculateTotalAmount() {
	fee := i.PaymentAmount.Mul(i.FeeRate).Floor()
	tax := fee.Mul(i.TaxRate).Floor()
	i.Fee = fee
	i.TaxAmount = tax
	i.TotalAmount = i.PaymentAmount.Add(fee).Add(tax)
}
