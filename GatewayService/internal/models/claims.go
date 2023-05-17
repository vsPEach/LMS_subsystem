package models

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	jwt.StandardClaims
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
