package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	idleTimeout    = 5 * time.Second
	maxHeaderBytes = 1 << 20
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(address string, port int) error {
	s.httpServer = s.makeHTTPServer(address, port)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) makeHTTPServer(address string, port int) *http.Server {
	return &http.Server{
		Addr:           fmt.Sprintf("%s:%d", address, port),
		Handler:        s.routes(),
		MaxHeaderBytes: maxHeaderBytes,
		IdleTimeout:    idleTimeout,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
	}
}

func (s *Server) routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
	})

	return r
}
