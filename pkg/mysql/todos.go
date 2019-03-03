package mysql

import (
	"database/sql"

	todo "github.com/KentaKudo/go-todo"
)

// TodoService provides an interface to handle todo.Todo.
type TodoService struct {
	db *sql.DB
}

// NewTodoService creates a new TodoService instance.
func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{db: db}
}

// Get returns a single todo.Todo instance.
func (ts *TodoService) Get(id int) (*todo.Todo, error) {
	var t todo.Todo
	if err := ts.db.QueryRow(`SELECT * FROM todos WHERE id = ?`, id).Scan(
		&t.ID,
		&t.Title,
	); err != nil {
		return nil, err
	}

	return &t, nil
}

// List returns a list of todo.Todo instances.
func (ts *TodoService) List() ([]todo.Todo, error) {
	rows, err := ts.db.Query(`SELECT * FROM todos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []todo.Todo{}
	for rows.Next() {
		var t todo.Todo
		if err := rows.Scan(&t.ID, &t.Title); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

// Create creates and stores a new todo.Todo instance.
func (ts *TodoService) Create(t *todo.Todo) error {
	r, err := ts.db.Exec(`INSERT INTO todos (id, title) VALUES (?, ?)`, t.ID, t.Title)
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return err
	}

	t.ID = int(id)
	return nil
}

// Delete deletes a stored todo.Todo instance.
func (ts *TodoService) Delete(id int) error {
	_, err := ts.db.Exec(`DELETE FROM todos WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}
