package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-cart/config"
	"shopping-cart/models"
)

type CreateOrderInput struct {
	CartID uint `json:"cart_id" binding:"required"`
}

func CreateOrder(c *gin.Context) {
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAny, _ := c.Get("user")
	user := userAny.(models.User)

	var cart models.Cart
	if err := config.DB.Where("id = ? AND user_id = ?", input.CartID, user.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	if cart.Status != "OPEN" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart is not open"})
		return
	}

	order := models.Order{
		CartID: cart.ID,
		UserID: user.ID,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	cart.Status = "ORDERED"
	config.DB.Save(&cart)

	c.JSON(http.StatusCreated, order)
}

func ListOrders(c *gin.Context) {
	userAny, _ := c.Get("user")
	user := userAny.(models.User)

	var orders []models.Order
	config.DB.Where("user_id = ?", user.ID).Find(&orders)

	c.JSON(http.StatusOK, orders)
}
