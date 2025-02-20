package models

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Name        string   `json:"name" gorm:"not null"`
	Description string   `json:"description" gorm:"not null"`
	Status      bool     `json:"status" gorm:"default:false"`
	TodoListID  uint     `json:"todo_list_id" gorm:"not null"`
	TodoList    TodoList `json:"-" gorm:"foreignKey:TodoListID"`
}
