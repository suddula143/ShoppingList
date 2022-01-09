package model

import "gorm.io/gorm"

//Category is to catagories the items and give their description
type Category struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Items       []Item
}
