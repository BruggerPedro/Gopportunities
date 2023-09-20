package router

import (
	"github.com/gin-gonic/gin"
)

func Inicialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8060")
}
