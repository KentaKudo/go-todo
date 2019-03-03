package todo

// Todo represents a todo item
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// TodoService represents a service object dealing with todos
type TodoService interface {
	Get(int) (*Todo, error)
	List() ([]Todo, error)
	Create(*Todo) error
	Delete(int) error
}
