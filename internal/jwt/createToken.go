package jwtpackage

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TO-DO func should get Role from database
func getRole(username string) string {
	if username == "me" {
		return "admin"
	} else {
		return "default"
	}
}

func CreateToken(username string) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "chatApp",
		"aud": getRole(username),
		"iat": time.Now().Unix(),
	})
	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return err.Error()
	}
	return tokenString
}
