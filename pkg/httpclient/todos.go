package httpclient

import (
	"net/http"

	todo "github.com/KentaKudo/go-todo"
)

type TodoService struct {
	cli *http.Client
}

func NewTodoService() *TodoService {
	return &TodoService{cli: http.DefaultClient}
}

func (ts *TodoService) Get(id int) (*todo.Todo, error) {
	// res, err := ts.cli.Get(fmt.Sprintf("http://localhost:8080/todos/%d", id))
	// if err != nil {
	// 	return nil, err
	// }
	// var t todo.Todo
	// if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
	// 	return nil, err
	// }

	return &todo.Todo{Title: "successfully get a todo"}, nil
}

func (ts *TodoService) List() ([]todo.Todo, error) {
	return []todo.Todo{todo.Todo{Title: "here we have a todo!"}}, nil
}

func (ts *TodoService) Create(t *todo.Todo) error {
	return nil
}

func (ts *TodoService) Delete(id int) error {
	return nil
}
