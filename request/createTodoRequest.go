package request

type CreateTodoRequest struct {
	Id          int    `validate:"required" json:"id"`
	Description string `validate:"required" json:"description"`
	Priority    string `validate:"required" json:"priority"`
}
