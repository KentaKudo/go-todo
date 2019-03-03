package httpserver

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/KentaKudo/go-todo/pkg/mysql"

	todo "github.com/KentaKudo/go-todo"
	"github.com/gorilla/mux"
)

// Server represents a server struct
type Server struct {
	Router *mux.Router

	TodoService todo.TodoService
}

// New returns a new Server instance
func New(db *sql.DB) *Server {
	ts := mysql.NewTodoService(db)
	return &Server{
		Router: mux.NewRouter(),

		TodoService: ts,
	}
}

// FIXME: Only for testing; remove if possible
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

// Run starts the server process
func (s *Server) Run(addr string) error {
	http.ListenAndServe(addr, s)
	return nil
}

type handlerFunc func(context.Context, http.ResponseWriter, *http.Request) error

// Error represents an error object.
type Error struct {
	Code int
	Err  error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func (fn handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if err := fn(ctx, w, r); err != nil {
		if ierr, ok := err.(*Error); ok {
			fail(w, ierr)
		} else {
			fail(w, &Error{Code: http.StatusInternalServerError, Err: err})
		}
	}
}
