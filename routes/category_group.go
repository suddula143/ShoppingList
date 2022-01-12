package routes

import "github.com/gin-gonic/gin"

func CategoryGroup(r *gin.Engine) {
	categoryGroup := r.Group("/shopping-list")
	{
		categoryGroup.POST("category", categoryController.AddCategory)
	}
}
