package model

import "gorm.io/gorm"

// ShoppingList gives the title for user shoppinglist
type ShoppingList struct {
	gorm.Model
	Title string `gorm:"not null"`
	Items []Item
	Users []User `gorm:"many2many:user_shopping_lists"`
}
