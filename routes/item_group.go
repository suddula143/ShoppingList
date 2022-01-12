package routes

import "github.com/gin-gonic/gin"

func ItemGroup(r *gin.Engine) {
	itemGroup := r.Group("/shopping-list")
	{
		itemGroup.POST("item", itemController.AddItem)
	}
}
