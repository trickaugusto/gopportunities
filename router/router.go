package router

import "github.com/gin-gonic/gin"

// When the first letter of a function's name is capitalized, it means that it is public and exported
func Initialize() {
	// Initialize the default Gin router
	// r is a ephemeral variable
	router := gin.Default()

	// Define a route handler for the GET /ping route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Run the server
	router.Run(":8080")
}
