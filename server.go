package skel

import (
	"context"
	"net/http"
)

// Server represents a server struct
type Server struct {
	router *http.ServeMux

	todoService TodoService
}

// New returns a new Server instance
func New(ts TodoService) *Server {
	return &Server{
		router: http.NewServeMux(),

		todoService: ts,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Run starts the server process
func (s *Server) Run(addr string) error {
	http.ListenAndServe(addr, s)
	return nil
}

type handlerFunc func(context.Context, http.ResponseWriter, *http.Request) error
