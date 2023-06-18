package controllers

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/Dmkk01/kanban-go-svelte/cmd/models"
	"github.com/Dmkk01/kanban-go-svelte/cmd/services"
	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

type loginCredentials struct {
	Email    string
	Password string
}

func Login(c echo.Context) error {
	var creds loginCredentials

	err := json.NewDecoder(c.Request().Body).Decode(&creds)
	if err != nil {
		return err
	}

	if creds.Email == "" {
		return echo.NewHTTPError(400, "Email is required")
	}
	if creds.Password == "" {
		return echo.NewHTTPError(400, "Password is required")
	}

	user, err := services.GetUserByEmail(creds.Email)
	if err != nil {
		return echo.NewHTTPError(500, "There is no user with that email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return echo.NewHTTPError(500, "Incorrect password")
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &models.Claims{
		Username: user.Username,
		Email:    user.Email,
		Id:       user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return echo.NewHTTPError(500, "Error generating token")
	}

	return c.JSON(200, struct {
		Token          string    `json:"token"`
		ExpirationTime time.Time `json:"expirationTime"`
	}{Token: tokenString, ExpirationTime: expirationTime})
}

type registerCredentials struct {
	Email    string
	Username string
	Password string
}

func Register(c echo.Context) error {
	var creds registerCredentials

	err := json.NewDecoder(c.Request().Body).Decode(&creds)
	if err != nil {
		return err
	}

	if creds.Email == "" {
		return echo.NewHTTPError(400, "Email is required")
	}
	if !govalidator.IsEmail(creds.Email) {
		return echo.NewHTTPError(400, "Invalid email")
	}

	if creds.Username == "" {
		return echo.NewHTTPError(400, "Username is required")
	}

	if len(creds.Username) < 3 {
		return echo.NewHTTPError(400, "Username must be at least 3 characters long")
	}

	//TODO better password validation
	if creds.Password == "" {
		return echo.NewHTTPError(400, "Password is required")
	}
	if len(creds.Password) < 6 {
		return echo.NewHTTPError(400, "Password must be at least 6 characters long")
	}

	ok := services.CreateUser(creds.Username, creds.Password, creds.Email)

	if !ok {
		return echo.NewHTTPError(500, "Error creating user")
	}

	return c.JSON(200, struct{ Status string }{Status: "OK"})
}

func RefreshToken(c echo.Context) error {
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
		return echo.NewHTTPError(400, "Bad request")
	}

	if !tkn.Valid {
		return echo.NewHTTPError(400, "Invalid token")
	}

	expirationTime := time.Now().Add(60 * time.Minute)

	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return echo.NewHTTPError(500, "Error generating token")
	}

	return c.JSON(200, struct {
		Token          string    `json:"token"`
		ExpirationTime time.Time `json:"expirationTime"`
	}{Token: tokenString, ExpirationTime: expirationTime})
}
