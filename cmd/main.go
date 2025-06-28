package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/i0li/super_shiharai_kun/internal/apperr"
	"github.com/i0li/super_shiharai_kun/internal/handler"
	"github.com/i0li/super_shiharai_kun/internal/infra/db"
	"github.com/i0li/super_shiharai_kun/internal/infra/logger"
	"github.com/i0li/super_shiharai_kun/internal/middleware"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"github.com/i0li/super_shiharai_kun/internal/usecase"
)

func main() {
	conn, err := db.NewGormPostgres()
	if err != nil {
		logger.L.Error(apperr.Wrap(err).DetailMessage())
	}

	userRepo := repository.NewUserRepository(conn)
	userUc := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUc)

	authUc := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUc)

	invoiceRepo := repository.NewInvoiceRepository(conn)
	invoiceUc := usecase.NewInvoiceUsecase(invoiceRepo)
	invoiceHandler := handler.NewInvoiceHandler(invoiceUc)

	r := gin.Default()
	r.POST("/users", userHandler.RegisterUser)
	r.POST("/auth/login", authHandler.Login)

	invoiceGroup := r.Group("/invoices")
	invoiceGroup.Use(middleware.JWTAuthMiddleware())
	{
		invoiceGroup.POST("", invoiceHandler.CreateInvoice)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
