package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TokenSecretKey = "airbnb-clone-backend/airbnb.com"

func GenerateToken(id int64, identifier, password string) (string, error) {
	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		identifier: identifier,
		"password": password,
		"id":       id,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	return tokenString.SignedString([]byte(TokenSecretKey))
}

func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(TokenSecretKey), nil
	})

	if err != nil {
		return -1, errors.New("invalid token")
	}

	validToken := parsedToken.Valid

	if !validToken {
		return -1, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return -1, errors.New("invalid token")
	}

	userId := int64(claims["id"].(float64))

	return userId, nil

}
