package repository

import (
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/model"
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	Create(invoice *model.Invoice) error
}

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) InvoiceRepository {
	return &invoiceRepository{db: db}
}

func (r *invoiceRepository) Create(invoice *model.Invoice) error {
	if err := r.db.Create(invoice).Error; err != nil {
		return apperr.Wrap(err)
	}
	return nil
}
