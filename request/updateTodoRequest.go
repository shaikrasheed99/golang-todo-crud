package request

type UpdateTodoRequest struct {
	Id          int    `validate:"required" json:"id"`
	Description string `validate:"required" json:"description"`
	Priority    string `validate:"required" json:"priority"`
}
