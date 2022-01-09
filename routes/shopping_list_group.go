package routes

import "github.com/gin-gonic/gin"

//ShoppingListGroup gives the path to connect to db
func ShoppingListGroup(r *gin.Engine) {
	shoppingListGroup := r.Group("/shopping-list")
	{
		shoppingListGroup.POST("create", shoppingListController.AddShoppingList)
	}
}
