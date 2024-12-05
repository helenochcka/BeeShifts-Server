package services

import (
	"BeeShifts-Server/internal/core/users"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService struct {
	SecretKey    string
	TokenExpTime int
}

func (as *AuthService) GenerateToken(userId int) (string, error) {
	payload := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Minute * time.Duration(as.TokenExpTime)).Unix(),
	}

	token, err := (jwt.NewWithClaims(jwt.SigningMethodHS256, payload)).SignedString([]byte(as.SecretKey))

	return token, err
}

func (as *AuthService) PayloadFromToken(tokenString string) (*users.TokenPayloadDTO, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(as.SecretKey), nil
	})

	if err != nil {
		return nil, as.mapJWTErrToUsersErr(err)
	}

	payload := token.Claims.(jwt.MapClaims)

	id := int(payload["id"].(float64))

	exp := int64(payload["exp"].(float64))

	tokenPayload := users.TokenPayloadDTO{Id: id, ExpiresAt: exp}

	return &tokenPayload, nil
}

func (as *AuthService) mapJWTErrToUsersErr(err error) error {
	switch {
	case errors.Is(err, jwt.ErrTokenExpired):
		return users.TokenExpired
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return users.TokenSignatureInvalid
	default:
		return err
	}
}
