package service

import (
	"errors"

	"github.com/alsolovyev/dummy-api/internal/entity"
)

type User struct {
}

func MakeUserService() *User {
	return &User{}
}

func (u *User) Create(_ *entity.User) error {
	return errors.New("Error occurred while creating user")
}
