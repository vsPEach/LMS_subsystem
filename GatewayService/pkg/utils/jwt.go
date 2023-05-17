package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/models"
	"time"
)

var Key = []byte("supersecretkey")

func GenerateJWT(email, password, role string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &models.JWTClaims{
		Email:    email,
		Password: password,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(Key)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(Key), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
