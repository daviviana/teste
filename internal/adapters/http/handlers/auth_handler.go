package handlers

import (
	"net/http"
	"teste/internal/ports"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	ports.AuthService
}

func NewAuthHandler(service ports.AuthService) ports.AuthHandler {
	return &AuthHandler{AuthService: service}
}

func (handler *AuthHandler) Login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Dados inválidos"})
	}

	token, err := handler.AuthService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func (handler *AuthHandler) RecoverPassword(c echo.Context) error {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Dados inválidos"})
	}

	token, err := handler.AuthService.GenerateRecoveryToken(req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"recovery_token": token})
}
