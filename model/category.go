package model

import "gorm.io/gorm"

//Category gives the Models
type Category struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Items       []Item
}
