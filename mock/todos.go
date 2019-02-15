package mock

import skel "github.com/KentaKudo/goapi-skel"

type todoService struct {
	GetTodosFn        func() ([]skel.Todo, error)
	GetTodosFnInvoked bool
}

// NewTodoService returns a todoService object which satisfies skel.todoService interface.
func NewTodoService(fn func() ([]skel.Todo, error)) *todoService {
	return &todoService{GetTodosFn: fn}
}

func (ts *todoService) GetTodos() ([]skel.Todo, error) {
	todos, err := ts.GetTodosFn()
	ts.GetTodosFnInvoked = true
	return todos, err
}
