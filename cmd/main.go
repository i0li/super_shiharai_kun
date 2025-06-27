package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/i0li/super_shiharai_kun/internal/db"
	"github.com/i0li/super_shiharai_kun/internal/handler"
	"github.com/i0li/super_shiharai_kun/internal/repository"
	"github.com/i0li/super_shiharai_kun/internal/usecase"
)

func main() {
	conn, err := db.NewGormPostgres()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	userRepo := repository.NewUserRepository(conn)
	userUc := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUc)

	authUc := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUc)

	r := gin.Default()
	r.POST("/user", userHandler.RegisterUser)
	r.POST("/auth/login", authHandler.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
