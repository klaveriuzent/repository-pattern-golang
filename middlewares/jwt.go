package middlewares

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var tokenHourLifespanStr = os.Getenv("TOKEN_HOUR_LIFESPAN")

func GenerateToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(parseDuration(tokenHourLifespanStr)).Unix()

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}

func parseDuration(durationStr string) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return time.Hour * 24
	}
	return duration
}
