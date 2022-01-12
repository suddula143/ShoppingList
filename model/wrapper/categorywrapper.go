package wrapper

import (
	"shoppinglist/model"

	"gorm.io/gorm"
)

type CategoryWrapper struct {
	DB *gorm.DB
}

func CreateCategoryWrapper(db *gorm.DB) *CategoryWrapper {
	return &CategoryWrapper{
		DB: db,
	}
}

func (s *CategoryWrapper) Insert(data interface{}) error {
	category := data.(*model.Category)

	return s.DB.Create(&category).Error
}

func (s *CategoryWrapper) Update(data interface{}) error {
	category := data.(*model.Category)

	return s.DB.Save(&category).Error
}

func (s *CategoryWrapper) Delete(data interface{}) error {
	category := data.(*model.Category)

	return s.DB.Delete(&category).Error
}
