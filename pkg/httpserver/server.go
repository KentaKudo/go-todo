package httpserver

import (
	"context"
	"net/http"

	app "github.com/KentaKudo/goapi-skel/pkg"
	"github.com/gorilla/mux"
)

// Server represents a server struct
type Server struct {
	router *mux.Router

	todoService app.TodoService
}

// New returns a new Server instance
func New(ts app.TodoService) *Server {
	return &Server{
		router: mux.NewRouter(),

		todoService: ts,
	}
}

// FIXME: Only for testing; remove if possible
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
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
