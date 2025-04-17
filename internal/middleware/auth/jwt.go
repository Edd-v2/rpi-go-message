package auth

import (
	"errors"
	"time"

	"github.com/Edd-v2/rpi-go-message/src/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.AppConfig.Auth.JWTSecret)

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Minute * time.Duration(config.AppConfig.Auth.TokenExpiration)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	userID, ok := claims["userId"].(string)
	if !ok {
		return "", errors.New("invalid user id in token")
	}

	return userID, nil
}
