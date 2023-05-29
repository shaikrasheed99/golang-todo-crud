package models

type Todo struct {
	Id          int    `gorm:"type:int;primaryKey;autoIncrement"`
	Description string `gorm:"type:varchar(255)"`
	Priority    string `gorm:"type:varchar(255)"`
}
