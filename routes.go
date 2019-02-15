package skel

import (
	"net/http"
)

// Routes defines and registers routes
func (s *Server) Routes() *Server {
	s.router.Handle("/todos", s.getTodos()).Methods(http.MethodGet)
	s.router.Handle("/todos", s.postTodo()).Methods(http.MethodPost)
	s.router.Handle("/todos/{id:[0-9]+}", s.getTodo()).Methods(http.MethodGet)
	s.router.Handle("/todos/{id:[0-9]+}", s.deleteTodo()).Methods(http.MethodDelete)

	return s
}
