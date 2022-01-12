package controller

import (
	"shoppinglist/model/wrapper"

	"gorm.io/gorm"
)

var (
	db                  *gorm.DB
	userWrapper         *wrapper.UserWrapper
	shoppingListWrapper *wrapper.ShoppingListWrapper
	categoryWrapper     *wrapper.CategoryWrapper
	itemWrapper         *wrapper.ItemWrapper
)

// InitializeController initializes the controller with a DB as parameter
func InitializeController(DB *gorm.DB) {
	db = DB
	userWrapper = wrapper.CreateUserWrapper(db)
	shoppingListWrapper = wrapper.CreateShoppingListWrapper(db)
	categoryWrapper = wrapper.CreateCategoryWrapper(db)
	itemWrapper = wrapper.CreateItemWrapper(db)
}
