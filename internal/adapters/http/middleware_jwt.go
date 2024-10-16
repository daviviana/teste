package http

import (
	"teste/internal/config"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	secret := config.AppConfig.JWTSecret

	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secret),
		ContextKey: "user",
	})
}

func ExtractUserFromToken(c echo.Context) (jwt.MapClaims, error) {
	user := c.Get("user").(jwt.MapClaims)

	if user != nil {
		return user, nil
	}
	return nil, echo.NewHTTPError(401, "Invalid JWT token")
}
