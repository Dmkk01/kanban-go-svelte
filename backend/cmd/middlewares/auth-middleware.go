package middlewares

import (
	"os"
	"strings"

	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.NewHTTPError(400, "No token provided")
		}

		token = strings.Replace(token, "Bearer ", "", 1)

		claims := &models.Claims{}

		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return echo.NewHTTPError(400, "Invalid signature token")
			}
			return echo.NewHTTPError(400, "Invalid token")
		}

		if !tkn.Valid {
			return echo.NewHTTPError(400, "Invalid token")
		}

		c.Set("claims", claims)

		return next(c)
	}
}
