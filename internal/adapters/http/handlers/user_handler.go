package handlers

import (
	"net/http"
	"teste/internal/adapters/db/models"
	"teste/internal/ports"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandler(service ports.UserService) ports.UserHandler {
	return &UserHandler{userService: service}
}

type User struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Address  Address `json:"address"`
}

type Address struct {
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	Number       uint   `json:"number"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zipcode"`
}

func (handler *UserHandler) RegisterUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	createdUser, err := handler.userService.RegisterUser(models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address: models.Address{
			Street:       user.Address.Street,
			Neighborhood: user.Address.Neighborhood,
			Number:       user.Address.Number,
			City:         user.Address.City,
			State:        user.Address.State,
			ZipCode:      user.Address.ZipCode,
		},
	},
		user.Password,
	)

	if err != nil && err.Error() == "UNIQUE constraint failed: users.email" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Email ja cadastrado"})
	}

	if err != nil && err.Error() == "CEP não encontrado" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "CEP não encontrado"})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, createdUser)
}

func (handler *UserHandler) ListUsers(c echo.Context) error {
	users, err := handler.userService.ListUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}
