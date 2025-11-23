package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"shopping-cart/config"
	"shopping-cart/routes"
)

func main() {
	// Initialize DB
	config.InitDB()

	r := gin.Default()
	r.Use(corsMiddleware())

	routes.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}

// Simple CORS middleware for local React dev
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
