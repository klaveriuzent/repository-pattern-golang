package middlewares

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}

	if err, ok := err.(*jwt.ValidationError); ok {
		if err.Errors&jwt.ValidationErrorExpired != 0 {
			return errors.New("token has expired")
		} else if err.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
			return errors.New("invalid token signature")
		} else {
			return errors.New("failed to validate token")
		}
	}

	if token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func getTokenFromRequestSample(context *gin.Context) (string, error) {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1], nil
	}
	return "", errors.New("invalid token format")
}
