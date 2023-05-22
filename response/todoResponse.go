package response

type TodoResponse struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}
