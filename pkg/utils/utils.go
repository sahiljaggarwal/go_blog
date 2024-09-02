package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("my-secret-key")

func GenerateJWT(id uint, email string, username string)(string, error){
	claims := jwt.MapClaims{
		"id":id,
		"email":email,
		"username":username,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
