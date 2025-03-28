package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string         `json:"email" gorm:"unique;not null"`
	Password      string         `json:"password" gorm:"not null"`
	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID"`
	TodoLists     []TodoList     `json:"todo_lists" gorm:"foreignKey:UserID"`
}
