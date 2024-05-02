package services

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type AuthService struct {
	SecretKey string
}

func (as *AuthService) GenerateToken(userId int) (string, error) {
	payload := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}

	token, err := (jwt.NewWithClaims(jwt.SigningMethodHS256, payload)).SignedString([]byte(as.SecretKey))

	return token, err
}

func (as *AuthService) PayloadFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(as.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	payload := token.Claims.(jwt.MapClaims)

	return payload, nil
}
