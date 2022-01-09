package model

import "gorm.io/gorm"

// Item in the shopping list gives the names of items in the list
type Item struct {
	gorm.Model
	Name           string `gorm:"not null"`
	Description    string
	IsBought       bool
	BoughtBy       string `gorm:"foreignkey:BoughtBy;association_foreignkey:UserID"`
	CategoryID     uint
	ShoppingListID uint
}
