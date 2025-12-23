package models

import "github.com/golang-jwt/jwt/v5"

type JWTCustomClaims struct {
	UserID uint `json:"user-id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}