package request

type UpdateTodoRequest struct {
	Description string `validate:"required" json:"description"`
	Priority    string `validate:"required" json:"priority"`
}
