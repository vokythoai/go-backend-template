package app

import (
	"log"
	handlers "qropen-backend/internal/adapters/handlers/http"
	"qropen-backend/internal/adapters/jwt"
	"qropen-backend/internal/adapters/middleware"
	"qropen-backend/internal/core/ports"
	"qropen-backend/internal/core/services"
	"qropen-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	app := &App{
		router: gin.Default(),
	}

	dbInstance, err := database.GetInstance("localhost", "5432", "thoaivo", "thoaivo", "qropen_development")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	repositories := ports.NewRepositories(dbInstance)
	jwtAdapter := jwt.NewJWTAdapter("your-secret-key")
	authService := services.NewAuthService(repositories, jwtAdapter)
	oauthService := services.NewOAuthService()
	authHandler := handlers.NewAuthHandler(authService, oauthService)

	app.router.POST("/login", authHandler.Login)
	app.router.GET("/logout", middleware.AuthMiddleware(authService), authHandler.Logout)
	app.router.GET("/protected", middleware.AuthMiddleware(authService), authHandler.Protected)

	return app
}

func (a *App) Run(addr string) error {
	return a.router.Run(addr)
}
