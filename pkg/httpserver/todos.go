package httpserver

import (
	"context"
	"net/http"
	"strconv"

	app "github.com/KentaKudo/goapi-skel/pkg"
	"github.com/gorilla/mux"
)

func (s *Server) getTodos() handlerFunc {
	type response struct {
		Todos []app.Todo `json:"todos"`
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
	type request app.Todo
	type response app.Todo
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) error {
		var req request
		if err := decode(r.Body, &req); err != nil {
			return err
		}

		if err := s.todoService.Create((*app.Todo)(&req)); err != nil {
			return err
		}

		return encode(w, response(req))
	}
}

func (s *Server) getTodo() handlerFunc {
	type response *app.Todo
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
