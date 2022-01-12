package model

import "gorm.io/gorm"

//Item is used to show  in ShoppingList
type Item struct {
	gorm.Model
	Name           string `gorm:"not null"`
	Description    string
	IsBought       bool
	BoughtBy       string `gorm:"foreignkey:BoughtBy;association_foreignkey:UserID"`
	CategoryID     uint
	ShoppingListID uint
}
