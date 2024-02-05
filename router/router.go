package router

import "github.com/gin-gonic/gin"

// When the first letter of a function's name is capitalized, it means that it is public and exported
func Initialize() {
	// Initiallize Router
	// r is a ephemeral variable
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
