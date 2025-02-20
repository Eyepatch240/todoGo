package models

import "gorm.io/gorm"

type TodoList struct {
	gorm.Model
	Name    string  `json:"name" gorm:"not null"`
	UserID  uint    `json:"user_id" gorm:"not null"`
	User    User    `json:"-" gorm:"foreignKey:UserID"`
	Entries []Entry `json:"entries" gorm:"foreignKey:TodoListID"`
}
