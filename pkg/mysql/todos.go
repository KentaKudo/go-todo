package mysql

import (
	app "github.com/KentaKudo/goapi-skel/pkg"
)

// TodoService provides an interface to handle skel.Todo.
type TodoService struct {
	client *Client
}

// NewTodoService creates a new TodoService instance.
func NewTodoService(client *Client) *TodoService {
	return &TodoService{client: client}
}

// Get returns a single skel.Todo instance.
func (ts *TodoService) Get(id int) (*app.Todo, error) {
	var todo app.Todo
	if err := ts.client.QueryRow(`SELECT * FROM todos WHERE id = ?`, id).Scan(
		&todo.ID,
		&todo.Title,
	); err != nil {
		return nil, err
	}

	return &todo, nil
}

// List returns a list of skel.Todo instances.
func (ts *TodoService) List() ([]app.Todo, error) {
	rows, err := ts.client.Query(`SELECT * FROM todos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []app.Todo{}
	for rows.Next() {
		var todo app.Todo
		if err := rows.Scan(&todo.ID, &todo.Title); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

// Create creates and stores a new skel.Todo instance.
func (ts *TodoService) Create(todo *app.Todo) error {
	r, err := ts.client.Exec(`INSERT INTO todos (id, title) VALUES (?, ?)`, todo.ID, todo.Title)
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return err
	}

	todo.ID = int(id)
	return nil
}

// Delete deletes a stored skel.Todo instance.
func (ts *TodoService) Delete(id int) error {
	_, err := ts.client.Exec(`DELETE FROM todos WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}
