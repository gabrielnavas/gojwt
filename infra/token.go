package infra

import (
	"errors"
	"gojwtexampleapi/config"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrMissingTokenValid = errors.New("missing token valid")
)

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrMissingTokenValid
		}
		return []byte(os.Getenv(config.ACCESS_SECRET)), nil
	})

	if err != nil {
		log.Println("jwt.Parse", err)
		return nil, ErrMissingTokenValid
	}

	if !token.Valid {
		log.Println(errors.New("token not valid"))
		return nil, ErrMissingTokenValid
	}
	return token, nil
}

func GetData(tokenStr string) (map[string]interface{}, error) {
	token, err := VerifyToken(tokenStr)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrMissingTokenValid
	}

	data := map[string]interface{}{}
	for key, value := range claims {
		data[key] = value
	}
	return data, nil
}
