package helpers

import (
	"time"
	"zunn/backend-api/config"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JTW_SECRET", "MY_VERY_SECRET_STRINGS123///"))

func GenerateToken(username string) string {
	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &jwt.RegisteredClaims{
		Subject:	username,
		ExpiresAt: 	jwt.NewNumericDate(expirationTime),
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	return token
}