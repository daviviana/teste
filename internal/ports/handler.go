package ports

import (
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	RegisterUser(c echo.Context) error
	ListUsers(c echo.Context) error
}

type AuthHandler interface {
	Login(c echo.Context) error
	RecoverPassword(c echo.Context) error
}
