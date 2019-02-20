package httpserver

import (
	"net/http"
)

// Routes defines and registers routes
func (s *Server) Routes() *Server {
	s.Router.Handle("/todos", s.getTodos()).Methods(http.MethodGet)
	s.Router.Handle("/todos", s.postTodo()).Methods(http.MethodPost)
	s.Router.Handle("/todos/{id:[0-9]+}", s.getTodo()).Methods(http.MethodGet)
	s.Router.Handle("/todos/{id:[0-9]+}", s.deleteTodo()).Methods(http.MethodDelete)

	return s
}
