package controller

import (
	"net/http"
	"shoppinglist/cache"
	"shoppinglist/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CategoryController struct{}

func (s *CategoryController) AddCategory(context *gin.Context) {

	status, _ := cache.AuthorizeSessionToken(context)
	if status != http.StatusOK {
		context.JSON(status, gin.H{
			"status":      status,
			"message":     "User unauthorized",
			"Category_id": 0,
		})
		return
	}
	//take the input
	inputCategory := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}{}

	if err := context.BindJSON(&inputCategory); err != nil {
		log.Error("Error parsing json input while creating the category", err.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Error is request body",
			"Category_id": 0,
		})
		return
	}

	var newCategory model.Category
	newCategory.Name = inputCategory.Name
	newCategory.Description = inputCategory.Description

	// user, _ := userWrapper.GetUser(string(resp.([]byte)))
	// newCategory.Users = make([]model.User, 0)
	// newCategory.Users = append(newCategory.Users, user)

	// create Category list
	err := categoryWrapper.Insert(&newCategory)
	if err != nil {
		log.Error(
			"Error creating Category list:",
			"Category Name:",
			inputCategory.Name,
			err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "Internal server error",
			"category_id": 0,
		})
		return
	}

	if newCategory.ID == 0 {
		log.Error(
			"Error creating Category list:",
			"Category Name:",
			inputCategory.Name,
		)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "Internal Server Error",
			"Category_id": 0,
		})
		return
	}

	//all checks passed
	log.Infof("Category successfully created Category Name:%d", newCategory.Name)
	context.JSON(http.StatusCreated, gin.H{
		"status":      http.StatusCreated,
		"message":     "Category successfully created",
		"Category Id": newCategory.ID,
	})
}
