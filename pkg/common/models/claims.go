package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID uint `json:"id" binding:"required"`
	jwt.StandardClaims
}
