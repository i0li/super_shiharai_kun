package server

import (
	"github.com/i0li/super_shiharai_kun/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(conn *gorm.DB) *gin.Engine {
	r := gin.Default()

	h := NewHandler(conn)

	r.POST("/users", h.User.RegisterUser)
	r.POST("/auth/login", h.Auth.Login)

	invoiceGroup := r.Group("/invoices")
	invoiceGroup.Use(middleware.JWTAuthMiddleware())
	{
		invoiceGroup.POST("", h.Invoice.CreateInvoice)
		invoiceGroup.GET("/payable", h.Invoice.FindPayableInvoicesInPeriod)
	}

	return r
}
