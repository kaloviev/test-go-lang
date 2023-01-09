package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"
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
		MaxHeaderBytes: maxHeaderBytes,
		IdleTimeout:    idleTimeout,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
	}
}
