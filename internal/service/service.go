package service

import "github.com/alsolovyev/dummy-api/internal/entity"

type UserService interface {
	Create(u *entity.User) error
}

type Service struct {
	User UserService
}

func New() *Service {
	return &Service{
		User: MakeUserService(),
	}
}
