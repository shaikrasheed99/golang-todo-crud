package models

type Todo struct {
	Id          int    `gorm:"type:int;primary_key"`
	Description string `gorm:"type:varchar(255)"`
	Priority    string `gorm:"type:varchar(255)"`
}

var Todos = []Todo{
	{Id: 1, Description: "Sleeping", Priority: "high"},
	{Id: 3, Description: "Reading", Priority: "medium"},
	{Id: 2, Description: "Playing", Priority: "low"},
}
