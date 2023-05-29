package request

type CreateTodoRequest struct {
	Description string `validate:"required" json:"description"`
	Priority    string `validate:"required" json:"priority"`
}
