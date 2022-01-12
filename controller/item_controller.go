package controller

import (
	"net/http"
	"shoppinglist/cache"
	"shoppinglist/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ItemController struct{}

func (s *ItemController) AddItem(context *gin.Context) {

	status, resp := cache.AuthorizeSessionToken(context)
	if status != http.StatusOK {
		context.JSON(status, gin.H{
			"status":  status,
			"message": "User unauthorized",
			"Item_id": 0,
		})
		return
	}
	//take the input
	inputItem := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		IsBought    bool   `json:"isbought"`
		// BoughtBy       string `json:"boughtby"`
		CategoryID     uint `json:"categoryid"`
		ShoppingListID uint `json:"shoppinglistid"`
	}{}

	if err := context.BindJSON(&inputItem); err != nil {
		log.Error("Error parsing json input while creating the item", err.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error is request body",
			"Item_id": 0,
		})
		return
	}

	var newItem model.Item
	newItem.Name = inputItem.Name
	newItem.Description = inputItem.Description
	newItem.IsBought = inputItem.IsBought
	newItem.CategoryID = inputItem.CategoryID
	newItem.ShoppingListID = inputItem.ShoppingListID

	//Getting userIdfrom response
	user, _ := userWrapper.GetUser(string(resp.([]byte)))
	newItem.BoughtBy = user.UserID

	// create Item list
	err := itemWrapper.Insert(&newItem)
	if err != nil {
		log.Error(
			"Error creating Item list:",
			"Item Name:",
			inputItem.Name,
			err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal server error",
			"item_id": 0,
		})
		return
	}

	if newItem.ID == 0 {
		log.Error(
			"Error creating Item list:",
			"Item Name:",
			inputItem.Name,
		)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
			"Item_id": 0,
		})
		return
	}

	//all checks passed
	log.Infof("Item successfully Added")
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Item successfully Added",
		"item_id": newItem.ID,
	})
}
