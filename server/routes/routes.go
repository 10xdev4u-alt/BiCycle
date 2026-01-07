package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/handlers"
)

// SetupRouter configures the API routes for the application.
func SetupRouter(router *gin.Engine) {
	// Group routes under /api/v1
	v1 := router.Group("/api/v1")
	{
		// Stand-related routes
		standRoutes := v1.Group("/stand")
		{
			standRoutes.POST("/scan", handlers.ScanStand)
		}

		// Cycle-related routes
		cycleRoutes := v1.Group("/cycles")
		{
			cycleRoutes.POST("/scan-book", handlers.ScanAndBookCycle)
			cycleRoutes.POST("/photo-pickup", handlers.UploadPickupPhoto)
		}

		// Other resource routes can be added here
		// e.g., cycleRoutes, userRoutes, etc.
	}
}
