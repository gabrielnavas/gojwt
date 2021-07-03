package infra

import (
	"gojwtexampleapi/config"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MakeTokenImpl struct{}

func NewMakeToken() *MakeTokenImpl {
	return &MakeTokenImpl{}
}

func (m MakeTokenImpl) Handle(data map[string]string, expiration int64) (token string, err error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true

	for key, value := range data {
		atClaims[key] = value
	}

	atClaims["exp"] = time.Now().UTC().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	keySecret := os.Getenv(config.ACCESS_SECRET)
	token, err = at.SignedString([]byte(keySecret))

	if err != nil {
		return "", err
	}
	return token, nil
}
