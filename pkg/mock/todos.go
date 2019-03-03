package mock

import todo "github.com/KentaKudo/go-todo"

// TodoService represents a mock object of todo.TodoService interface.
type TodoService struct {
	GetFn        func(int) (*todo.Todo, error)
	GetFnInvoked bool

	ListFn        func() ([]todo.Todo, error)
	ListFnInvoked bool

	CreateFn        func(*todo.Todo) error
	CreateFnInvoked bool

	DeleteFn        func(int) error
	DeleteFnInvoked bool
}

// NewTodoService creates a new mock TodoService instance.
func NewTodoService() *TodoService {
	return &TodoService{
		GetFn:    func(_ int) (*todo.Todo, error) { return &todo.Todo{}, nil },
		ListFn:   func() ([]todo.Todo, error) { return []todo.Todo{}, nil },
		CreateFn: func(_ *todo.Todo) error { return nil },
		DeleteFn: func(_ int) error { return nil },
	}
}

// Get invokes the GetFn method.
func (ts *TodoService) Get(id int) (*todo.Todo, error) {
	todo, err := ts.GetFn(id)
	ts.GetFnInvoked = true
	return todo, err
}

// List invokes the ListFn method.
func (ts *TodoService) List() ([]todo.Todo, error) {
	todos, err := ts.ListFn()
	ts.ListFnInvoked = true
	return todos, err
}

// Create invokes the CreateFn method.
func (ts *TodoService) Create(todo *todo.Todo) error {
	err := ts.CreateFn(todo)
	ts.CreateFnInvoked = true
	return err
}

// Delete invokes the DeleteFn method.
func (ts *TodoService) Delete(id int) error {
	err := ts.DeleteFn(id)
	ts.DeleteFnInvoked = true
	return err
}
