package model

import "gorm.io/gorm"

// ShoppingList  this specifies  title for the shoppinlist for each user
type ShoppingList struct {
	gorm.Model
	Title string `gorm:"not null"`
	Items []Item
	Users []User `gorm:"many2many:user_shopping_lists;association_foreignkey:UserID"`
}
