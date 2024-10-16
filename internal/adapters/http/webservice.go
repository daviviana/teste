package http

import (
	"net/http"
	"teste/internal/adapters"
	"teste/internal/adapters/db/repositories"
	"teste/internal/adapters/http/handlers"
	"teste/internal/adapters/zip"
	"teste/internal/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewWebService() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())
	registerRoutes(e)

	return e
}

func registerRoutes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	UserRoutes(e)
	AuthRoutes(e)
}

func UserRoutes(e *echo.Echo) {
	userRepo := repositories.NewUserRepository("test.db")
	zipAdapter := zip.NewZipAdapter()
	userService := domain.NewUserService(userRepo, zipAdapter)
	userHandler := handlers.NewUserHandler(userService)

	userRoutes := e.Group("/users")
	userRoutes.POST("", userHandler.RegisterUser)

	userRoutes.Use(JWTMiddleware())
	userRoutes.GET("", userHandler.ListUsers)
}

func AuthRoutes(e *echo.Echo) {
	userRepo := repositories.NewUserRepository("test.db")
	emailAdapter := adapters.NewEmailAdapter()
	authService := domain.NewLoginService(userRepo, emailAdapter)
	authHandler := handlers.NewAuthHandler(authService)

	authRoutes := e.Group("")
	authRoutes.POST("/login", authHandler.Login)
	authRoutes.POST("/recover", authHandler.RecoverPassword)
}
