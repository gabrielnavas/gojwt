package usecase

import (
	"fmt"
	"gojwtexampleapi/config"
	"gojwtexampleapi/entity"
	"gojwtexampleapi/infra"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type GetDataFromToken func(string) (map[string]interface{}, error)

func ExtractTokenMetadata(getData GetDataFromToken) func(tokenStr string) (*entity.User, error) {
	return func(tokenStr string) (*entity.User, error) {
		data, err := getData(tokenStr)
		if err != nil {
			return nil, infra.ErrMissingTokenValid
		}

		user_id, err := strconv.ParseUint(fmt.Sprintf("%.f", data["user_id"]), 10, 64)
		if err != nil {
			return nil, infra.ErrMissingTokenValid
		}

		user_name, ok := data["user_name"].(string)
		if !ok {
			return nil, infra.ErrMissingTokenValid
		}
		return &entity.User{
			ID:   user_id,
			Name: user_name,
		}, nil
	}
}
func MakeToken(user entity.User) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["user_name"] = user.Name
	atClaims["exp"] = time.Now().UTC().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	keySecret := os.Getenv(config.ACCESS_SECRET)
	token, err := at.SignedString([]byte(keySecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
