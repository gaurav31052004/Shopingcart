package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-cart/config"
	"shopping-cart/models"
)

type CreateItemInput struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status"`
}

func CreateItem(c *gin.Context) {
	var input CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.Item{
		Name:   input.Name,
		Status: input.Status,
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func ListItems(c *gin.Context) {
	var items []models.Item
	config.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}
