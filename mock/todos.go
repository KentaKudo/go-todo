package mock

import skel "github.com/KentaKudo/goapi-skel"

// TodoService represents a mock object of skel.TodoService interface.
type TodoService struct {
	GetFn        func(int) (*skel.Todo, error)
	GetFnInvoked bool

	ListFn        func() ([]skel.Todo, error)
	ListFnInvoked bool

	CreateFn        func(*skel.Todo) error
	CreateFnInvoked bool

	DeleteFn        func(int) error
	DeleteFnInvoked bool
}

// NewTodoService creates a new mock TodoService instance.
func NewTodoService() *TodoService {
	return &TodoService{
		GetFn:    func(_ int) (*skel.Todo, error) { return &skel.Todo{}, nil },
		ListFn:   func() ([]skel.Todo, error) { return []skel.Todo{}, nil },
		CreateFn: func(_ *skel.Todo) error { return nil },
		DeleteFn: func(_ int) error { return nil },
	}
}

// Get invokes the GetFn method.
func (ts *TodoService) Get(id int) (*skel.Todo, error) {
	todo, err := ts.GetFn(id)
	ts.GetFnInvoked = true
	return todo, err
}

// List invokes the ListFn method.
func (ts *TodoService) List() ([]skel.Todo, error) {
	todos, err := ts.ListFn()
	ts.ListFnInvoked = true
	return todos, err
}

// Create invokes the CreateFn method.
func (ts *TodoService) Create(todo *skel.Todo) error {
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
