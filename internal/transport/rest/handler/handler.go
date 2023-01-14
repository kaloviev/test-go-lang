package handler

import (
	"github.com/alsolovyev/dummy-api/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type UserService interface {
	Create(u *entity.User) error
}

type Handler struct {
	UserService
}

func New(u UserService) *Handler {
	return &Handler{
		UserService: u,
	}
}

func (h *Handler) MakeHandler() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", h.PingHandler)

		r.Post("/user/create", h.CreateUser)
	})

	return r
}
