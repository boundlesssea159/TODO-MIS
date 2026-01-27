package dto

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
}

type CreateTodoResponse struct {
	ID int `json:"id"`
}
