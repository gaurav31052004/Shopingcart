package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-cart/config"
	"shopping-cart/models"
)

type AddToCartInput struct {
	ItemID uint `json:"item_id" binding:"required"`
}

func AddToCart(c *gin.Context) {
	var input AddToCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAny, _ := c.Get("user")
	user := userAny.(models.User)

	// Find existing open cart or create new
	var cart models.Cart
	if err := config.DB.
		Where("user_id = ? AND status = ?", user.ID, "OPEN").
		First(&cart).Error; err != nil {

		cart = models.Cart{
			UserID: user.ID,
			Name:   "Cart for " + user.Username,
			Status: "OPEN",
		}

		if err := config.DB.Create(&cart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create cart"})
			return
		}
	}

	cartItem := models.CartItem{
		CartID: cart.ID,
		ItemID: input.ItemID,
	}

	if err := config.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add item to cart"})
		return
	}

	// load items
	config.DB.Preload("Items").First(&cart, cart.ID)

	c.JSON(http.StatusOK, cart)
}

func GetMyCart(c *gin.Context) {
	userAny, _ := c.Get("user")
	user := userAny.(models.User)

	var cart models.Cart
	if err := config.DB.Preload("Items").
		Where("user_id = ? AND status = ?", user.ID, "OPEN").
		First(&cart).Error; err != nil {

		c.JSON(http.StatusOK, gin.H{"message": "no open cart", "cart": nil})
		return
	}

	c.JSON(http.StatusOK, cart)
}
