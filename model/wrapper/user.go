package wrapper

import (
	"errors"
	"shoppinglist/model"

	"gorm.io/gorm"
)

//UserWrapper  hides the details of the user in the database
type UserWrapper struct {
	DB *gorm.DB
}

//CreateUserWrapper hides the details of the users
func CreateUserWrapper(db *gorm.DB) *UserWrapper {
	return &UserWrapper{
		DB: db,
	}
}

// Insert is used to insert the data of the user
func (u *UserWrapper) Insert(data interface{}) error {
	user := data.(*model.User)

	return u.DB.Create(&user).Error
}

//Update is used to update the user data in the wrapper
func (u *UserWrapper) Update(data interface{}) error {
	user := data.(*model.User)

	return u.DB.Save(&user).Error
}

//GetUser   is to get the GetUser data belonging to the user
func (u *UserWrapper) GetUser(userID string) (model.User, error) {
	var user model.User

	err := u.DB.Where(&model.User{UserID: userID}).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

//GetAllUser is to get the data of the all users
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

//UserAlreadyExists exits verify whether the user exits are not
func (u *UserWrapper) UserAlreadyExists(userID string) bool {
	var user model.User
	err := u.DB.Where(&model.User{UserID: userID}).First(&user).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

//GetShoppingLists gets the shoppinglist of the user
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
