package wrapper

import (
	"shoppinglist/model"

	"gorm.io/gorm"
)

type ItemWrapper struct {
	DB *gorm.DB
}

func CreateItemWrapper(db *gorm.DB) *ItemWrapper {
	return &ItemWrapper{
		DB: db,
	}
}

func (s *ItemWrapper) Insert(data interface{}) error {
	item := data.(*model.Item)

	return s.DB.Create(&item).Error
}

func (s *ItemWrapper) Update(data interface{}) error {
	item := data.(*model.Item)

	return s.DB.Save(&item).Error
}

func (s *ItemWrapper) Delete(data interface{}) error {
	item := data.(*model.Item)

	return s.DB.Delete(&item).Error
}
