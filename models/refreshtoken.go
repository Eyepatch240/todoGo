package models

import "gorm.io/gorm"

type RefreshToken struct {
	gorm.Model
	Token  string `json:"token" gorm:"unique;not null"`
	UserID uint   `json:"user_id" gorm:"not null"`
}
