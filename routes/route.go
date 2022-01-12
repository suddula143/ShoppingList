package routes

import (
	"shoppinglist/controller"

	"github.com/gin-gonic/gin"
)

var (
	signupController       = new(controller.SignupController)
	userController         = new(controller.UserController)
	shoppingListController = new(controller.ShoppingListController)
	categoryController     = new(controller.CategoryController)
	itemController         = new(controller.ItemController)
)

//InitRoutes give connection to database
func InitRoutes(r *gin.Engine) {
	UserGroup(r)
	ShoppingListGroup(r)
	CategoryGroup(r)
	ItemGroup(r)
}
