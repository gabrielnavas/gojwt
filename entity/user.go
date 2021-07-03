package entity

import "github.com/google/uuid"

type User struct {
	ID   string
	Name string
}

func NewUser(id, name string) *User {
	if id == "" {
		id = uuid.New().String()
	}
	u := &User{
		ID:   id,
		Name: name,
	}
	return u
}

func (u User) Validate() error {
	return nil
}
