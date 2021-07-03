package usecase

import (
	"fmt"
	"gojwtexampleapi/entity"
	"time"
)

type MakeTokenInfra interface {
	Handle(data map[string]string, expiration int64) (token string, err error)
}

type CreateToken interface {
	Handle(id, name string) (string, error)
}

type createTokenImpl struct {
	makeToken MakeTokenInfra
}

func NewCreateTokenImpl(makeToken MakeTokenInfra) *createTokenImpl {
	return &createTokenImpl{makeToken}
}

// type name to data inside token
const (
	UserId   = "user_id"
	UserName = "user_name"
)

func (c createTokenImpl) Handle(id, name string) (string, error) {
	u := entity.NewUser(id, name)
	err := u.Validate()
	if err != nil {
		return "", err
	}

	// ALERT!
	// fictitious user on under, you need find the user in repository
	// user, err := c.repository.findUserById(id)
	// if err != nil {
	// 	return nil, userNotFound
	// }

	data := map[string]string{
		UserId:   fmt.Sprint(u.ID),
		UserName: u.Name,
	}
	exp := time.Now().UTC().Add(time.Minute * 15).Unix()
	token, err := c.makeToken.Handle(data, exp)
	if err != nil {
		return "", err
	}
	return token, nil
}
