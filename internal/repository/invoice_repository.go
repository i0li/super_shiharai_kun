package repository

import (
	"time"

	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/domain"
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	Create(invoice *domain.Invoice) error
	FindByPaymentDueDatePeriod(
		userID int64,
		startDate, endDate time.Time,
	) ([]*domain.Invoice, error)
}

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) InvoiceRepository {
	return &invoiceRepository{db: db}
}

func (r *invoiceRepository) Create(invoice *domain.Invoice) error {
	if err := r.db.Create(invoice).Error; err != nil {
		return apperr.Wrap(err)
	}
	return nil
}

func (r *invoiceRepository) FindByPaymentDueDatePeriod(
	userID int64,
	startDate, endDate time.Time,
) ([]*domain.Invoice, error) {
	var invoices []*domain.Invoice
	err := r.db.
		Where("user_id = ?", userID).
		Where("payment_due_date BETWEEN ? AND ?", startDate, endDate).
		Order("payment_due_date DESC").
		Find(&invoices).Error
	if err != nil {
		return nil, apperr.Wrap(err)
	}

	return invoices, nil
}
