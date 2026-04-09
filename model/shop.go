package model

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Image       string `json:"image"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Status      int    `json:"status" gorm:"default:1"`
	UserID      uint   `json:"user_id"`
}
