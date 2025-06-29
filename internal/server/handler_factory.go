package server

import (
	"github.com/i0li/super_shiharai_kun/internal/handler"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"github.com/i0li/super_shiharai_kun/internal/usecase"

	"gorm.io/gorm"
)

type Handlers struct {
	User    *handler.UserHandler
	Auth    *handler.AuthHandler
	Invoice *handler.InvoiceHandler
}

func NewHandler(db *gorm.DB) *Handlers {
	userRepo := repository.NewUserRepository(db)

	return &Handlers{
		User: handler.NewUserHandler(usecase.NewUserUsecase(userRepo)),
		Auth: handler.NewAuthHandler(usecase.NewAuthUsecase(userRepo)),
		Invoice: handler.NewInvoiceHandler(
			usecase.NewInvoiceUsecase(repository.NewInvoiceRepository(db)),
		),
	}
}
