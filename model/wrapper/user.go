package wrapper

import (
	"errors"
	"shoppinglist/model"

	"gorm.io/gorm"
)

//UserWrapper hide details of the User in DB
type UserWrapper struct {
	DB *gorm.DB
}

//CreateUserWrapper hide the details of User
func CreateUserWrapper(db *gorm.DB) *UserWrapper {
	return &UserWrapper{
		DB: db,
	}
}

//Insert is used to Insert the User data in DB
func (u *UserWrapper) Insert(data interface{}) error {
	user := data.(*model.User)

	return u.DB.Create(&user).Error
}

//Update is used to Update User data in DB
func (u *UserWrapper) Update(data interface{}) error {
	user := data.(*model.User)

	return u.DB.Save(&user).Error
}

//GetUser is used to GetUser data from ShoppingList
func (u *UserWrapper) GetUser(userID string) (model.User, error) {
	var user model.User

	err := u.DB.Where(&model.User{UserID: userID}).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

//GetAllUser is used to getall user data
func (u *UserWrapper) GetAllUser() []string {
	var users []model.User

	userIDs := make([]string, 0)

	err := u.DB.Find(&users).Error
	if err != nil {
		return userIDs
	}

	for _, user := range users {
		userIDs = append(userIDs, user.UserID)
	}

	return userIDs
}

//UserAlreadyExists gives the user exists details
func (u *UserWrapper) UserAlreadyExists(userID string) bool {
	var user model.User
	err := u.DB.Where(&model.User{UserID: userID}).First(&user).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

//GetShoppingLists gives the shoppinglists user details
func (u *UserWrapper) GetShoppingLists(userID string) []model.ShoppingList {
	user, err := u.GetUser(userID)
	if err != nil {
		return nil
	}

	u.DB.Preload("ShoppingLists").Find(&user)
	for i := 0; i < len(user.ShoppingLists); i++ {
		u.DB.Preload("Users").Find(&user.ShoppingLists[i])
		u.DB.Preload("Items").Find(&user.ShoppingLists[i])
	}

	return user.ShoppingLists
}
