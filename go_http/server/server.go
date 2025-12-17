package server

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	port         int
	Router       *Router
	middelwares  []Middleware
	errorHandler ErrorHandler
	server       *http.Server
}

// ServerOption defines a function type for configuring the Server.
type ServerOption func(*Server)

// NewServer creates a new Server instance with the specified port.
func NewServer(port int) *Server {
	s := &Server{
		port:        port,
		middelwares: []Middleware{},
	}
	s.Router = NewRouter(s)
	return s
}

// ServeHTTP implements the http.Handler interface for the Server.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

// Start starts the HTTP server on the specified port.
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.port)
	s.server = &http.Server{
		Addr:    addr,
		Handler: s,
	}

	fmt.Printf("Starting server on port %d\n", s.port)
	return s.server.ListenAndServe()
}

// shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

// get the router
func (s *Server) GetRouter() *Router {
	return s.Router
}

//error handling

type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

func withErrorHandler(handler ErrorHandler) ServerOption {
	return func(s *Server) {
		s.errorHandler = handler
	}
}

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
