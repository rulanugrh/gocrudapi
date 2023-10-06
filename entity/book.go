package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name       string `json:"name" validate:"required"`
	Year       int    `json:"year" validate:"required"`
	Price      int    `json:"price" validate:"required"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

type Category struct {
	gorm.Model
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
