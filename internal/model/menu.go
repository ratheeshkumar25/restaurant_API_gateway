package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Category  string  `json:"category" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
	FoodImage string  `json:"foodImage" validate:"required"`
	Duration  string  `json:"duration" validate:"required"`
}
