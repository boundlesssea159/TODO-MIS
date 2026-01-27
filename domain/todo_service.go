package domain

type Todo struct {
	repository TodoRepository
}

func NewTodo(repository TodoRepository) *Todo {
	return &Todo{
		repository: repository,
	}
}

func (todo *Todo) Create(title string, description string) (int, error) {
	return todo.repository.Create(title, description)
}
