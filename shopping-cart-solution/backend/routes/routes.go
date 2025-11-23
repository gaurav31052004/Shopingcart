package routes

import (
	"github.com/gin-gonic/gin"

	"shopping-cart/controllers"
	"shopping-cart/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	// Users
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.ListUsers)
	r.POST("/users/login", controllers.Login)

	// Items
	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.ListItems)

	// Authenticated routes
	auth := r.Group("/")
	auth.Use(middleware.AuthRequired())

	auth.POST("/carts", controllers.AddToCart)
	auth.GET("/carts", controllers.GetMyCart)

	auth.POST("/orders", controllers.CreateOrder)
	auth.GET("/orders", controllers.ListOrders)
}
