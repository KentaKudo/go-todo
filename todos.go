package skel

import (
	"context"
	"encoding/json"
	"net/http"
)

// Todo represents a todo item
type Todo struct {
	Title string `json:"title"`
}

// TodoService represents a service object dealing with todos
type TodoService interface {
	GetTodos() ([]Todo, error)
}

func (s *Server) getTodos() handlerFunc {
	type response struct {
		Todos []Todo `json:"todos"`
	}
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) error {
		todos, err := s.todoService.GetTodos()
		if err != nil {
			return nil
		}
		return json.NewEncoder(w).Encode(response{Todos: todos})
	}
}

func (s *Server) postTodo() handlerFunc {
	type request Todo
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) error {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(req)
	}
}
