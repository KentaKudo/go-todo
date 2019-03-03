package httpserver

import (
	"context"
	"net/http"
	"strconv"

	todo "github.com/KentaKudo/go-todo"
	"github.com/gorilla/mux"
)

func (s *Server) getTodos() handlerFunc {
	type response struct {
		Todos []todo.Todo `json:"todos"`
	}
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) error {
		todos, err := s.TodoService.List()
		if err != nil {
			return nil
		}
		return encode(w, response{Todos: todos})
	}
}

func (s *Server) postTodo() handlerFunc {
	type request todo.Todo
	type response todo.Todo
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) error {
		var req request
		if err := decode(r.Body, &req); err != nil {
			return err
		}

		if err := s.TodoService.Create((*todo.Todo)(&req)); err != nil {
			return err
		}

		return encode(w, response(req))
	}
}

func (s *Server) getTodo() handlerFunc {
	type response *todo.Todo
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) error {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return err
		}

		todo, err := s.TodoService.Get(id)
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

		return s.TodoService.Delete(id)
	}
}
