package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
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
