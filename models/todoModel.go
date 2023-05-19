package models

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

var Todos = []Todo{
	{Id: 1, Description: "Sleeping", Priority: "high"},
	{Id: 3, Description: "Reading", Priority: "medium"},
	{Id: 2, Description: "Playing", Priority: "low"},
}
