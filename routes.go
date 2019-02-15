package skel

import (
	"context"
	"fmt"
	"net/http"
)

type endpoint map[string]handlerFunc

func (e endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn, ok := e[r.Method]
	if !ok {
		fn = methodNotAllowedHandler
	}

	if err := fn(context.Background(), w, r); err != nil {
		// TODO
	}
}

func methodNotAllowedHandler(_ context.Context, w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("405 method not found")
}

// Routes defines and registers routes
func (s *Server) Routes() *Server {
	todos := endpoint{
		http.MethodGet:  s.getTodos(),
		http.MethodPost: s.postTodo(),
	}
	s.router.Handle("/todos", todos)

	return s
}
