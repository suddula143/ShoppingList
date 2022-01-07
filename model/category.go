package model

import "gorm.io/gorm"

//Catgory is used to ShoppingList

type Category struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Items       []Item
}
