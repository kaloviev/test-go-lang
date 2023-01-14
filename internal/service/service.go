package service

import "github.com/alsolovyev/dummy-api/internal/transport/rest/handler"

type Service struct {
	User handler.UserService
}

func New() *Service {
	return &Service{
		User: MakeUserService(),
	}
}
