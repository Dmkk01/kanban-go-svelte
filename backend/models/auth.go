package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       int    `json:"id"`
	jwt.RegisteredClaims
}
