package usecase

import (
	"errors"
	"gojwtexampleapi/entity"
)

var (
	ErrVerifyTokenDataNotFound = errors.New("error on verify token")
)

type GetTokenInfra interface {
	Handle(string) (map[string]string, error)
}

type VerifyToken interface {
	Handle(token string) (*entity.User, error)
}

type VerifyTokenImpl struct {
	token GetTokenInfra
}

func NewVerifyToken(token GetTokenInfra) *VerifyTokenImpl {
	return &VerifyTokenImpl{token}
}

func (v VerifyTokenImpl) Handle(tokenStr string) (*entity.User, error) {
	data, err := v.token.Handle(tokenStr)
	if err != nil {
		return nil, ErrVerifyTokenDataNotFound
	}

	user_id, ok := data[UserId]
	if err != nil {
		return nil, ErrVerifyTokenDataNotFound
	}

	user_name, ok := data[UserName]
	if !ok {
		return nil, ErrVerifyTokenDataNotFound
	}

	u := &entity.User{
		ID:   user_id,
		Name: user_name,
	}

	return u, nil
}
