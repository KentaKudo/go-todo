package mysql

import (
	"database/sql"

	skel "github.com/KentaKudo/goapi-skel"
)

// TodoService provides an interface to handle skel.Todo.
type TodoService struct {
	db *sql.DB
}

// NewTodoService creates a new TodoService instance.
func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{db: db}
}

// Get returns a single skel.Todo instance.
func (ts *TodoService) Get(id int) (*skel.Todo, error) {
	var todo skel.Todo
	if err := ts.db.QueryRow(`SELECT * FROM todos WHERE id = ?`, id).Scan(
		&todo.ID,
		&todo.Title,
	); err != nil {
		return nil, err
	}

	return &todo, nil
}

// List returns a list of skel.Todo instances.
func (ts *TodoService) List() ([]skel.Todo, error) {
	rows, err := ts.db.Query(`SELECT * FROM todos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []skel.Todo{}
	for rows.Next() {
		var todo skel.Todo
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
func (ts *TodoService) Create(todo *skel.Todo) error {
	r, err := ts.db.Exec(`INSERT INTO todos (id, title) VALUES (?, ?)`, todo.ID, todo.Title)
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
	_, err := ts.db.Exec(`DELETE FROM todos WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}
