package skel

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Todo represents a todo item
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// TodoService represents a service object dealing with todos
type TodoService interface {
	Get(int) (*Todo, error)
	List() ([]Todo, error)
	Create(*Todo) error
	Delete(int) error
}

func (s *Server) getTodos() handlerFunc {
	type response struct {
		Todos []Todo `json:"todos"`
	}
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) error {
		todos, err := s.todoService.List()
		if err != nil {
			return nil
		}
		return encode(w, response{Todos: todos})
	}
}

func (s *Server) postTodo() handlerFunc {
	type request Todo
	type response Todo
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) error {
		var req request
		if err := decode(r.Body, &req); err != nil {
			return err
		}

		if err := s.todoService.Create((*Todo)(&req)); err != nil {
			return err
		}

		return encode(w, response(req))
	}
}

func (s *Server) getTodo() handlerFunc {
	type response *Todo
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) error {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return err
		}

		todo, err := s.todoService.Get(id)
		if err != nil {
			return err
		}

		return encode(w, response(todo))
	}
}

func (s *Server) deleteTodo() handlerFunc {
	type response struct {
		Result string `json:"result"`
	}
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) error {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return err
		}

		return s.todoService.Delete(id)
	}
}
