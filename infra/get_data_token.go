package infra

import (
	"errors"
	"gojwtexampleapi/config"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrMissingTokenValid = errors.New("missing token valid")
	ErrDataType          = errors.New("missing data type")
)

type GetDataInfraImpl struct{}

func NewGetDataInfra() *GetDataInfraImpl {
	return &GetDataInfraImpl{}
}

func (gd GetDataInfraImpl) Handle(tokenStr string) (map[string]string, error) {
	token, err := gd.verifyToken(tokenStr)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrMissingTokenValid
	}

	// try it better :(
	data := map[string]string{}
	for key, value := range claims {
		str, ok := value.(string)
		if ok {
			data[key] = str
		} else {
			continue
		}
	}
	return data, nil
}

func (gd GetDataInfraImpl) verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := getToken(tokenString)
	if err != nil {
		return nil, ErrMissingTokenValid
	}

	if !token.Valid {
		return nil, ErrMissingTokenValid
	}
	return token, nil
}

func getToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrMissingTokenValid
		}
		return []byte(os.Getenv(config.ACCESS_SECRET)), nil
	})
}
