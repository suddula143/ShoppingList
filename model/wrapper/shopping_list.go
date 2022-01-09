package wrapper

import (
	"shoppinglist/model"

	"gorm.io/gorm"
)

//ShoppingListWrapper connects to the database
type ShoppingListWrapper struct {
	DB *gorm.DB
}

//CreateShoppingListWrapper creates shopping list for each user
func CreateShoppingListWrapper(db *gorm.DB) *ShoppingListWrapper {
	return &ShoppingListWrapper{
		DB: db,
	}
}

//Insert shopping list inserts to the usershoppinglist
func (s *ShoppingListWrapper) Insert(data interface{}) error {
	shoppingList := data.(*model.ShoppingList)

	return s.DB.Create(&shoppingList).Error
}

//Update shoppinglist updates the list for each user
func (s *ShoppingListWrapper) Update(data interface{}) error {
	shoppingList := data.(*model.ShoppingList)

	return s.DB.Save(&shoppingList).Error
}

//Delete delets the shoppinglist for user
func (s *ShoppingListWrapper) Delete(data interface{}) error {
	shoppingList := data.(*model.ShoppingList)

	return s.DB.Delete(&shoppingList).Error
}
