package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func getSecretKey() ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("could not load .env file: " + err.Error())
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("JWT_SECRET_KEY environment variable is not set")
	}
	return []byte(secretKey), nil
}

func GenerateJWT(userId int64, email string) (string, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(secretKey)
}

func ValidateJWT(tokenString string) (int64, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return 0, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := token.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token")
	}

	userId := int64(claims["userId"].(float64))

	return userId, nil
}