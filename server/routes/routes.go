package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/api"
)

// SetupRouter configures the API routes for the application.
func SetupRouter(router *gin.Engine, api *api.API) {
	// Group routes under /api/v1
	v1 := router.Group("/api/v1")
	{
		// Stand-related routes
		standRoutes := v1.Group("/stand")
		{
			standRoutes.POST("/scan", api.ScanStand)
		}

		// Cycle-related routes
		cycleRoutes := v1.Group("/cycles")
		{
			cycleRoutes.POST("/scan-book", api.ScanAndBookCycle)
			cycleRoutes.POST("/photo-pickup", api.UploadPickupPhoto)
		}

		// Guard-related routes
		guardRoutes := v1.Group("/guard")
		{
			guardRoutes.POST("/scan", api.ScanTicket)
			guardRoutes.POST("/approve-pickup", api.ApprovePickup)
			guardRoutes.POST("/approve-return", api.ApproveReturn)
		}
	
		// Auth routes
		authRoutes := router.Group("/auth")
		{
			authRoutes.GET("/google", api.GoogleLogin)
			authRoutes.GET("/google/callback", api.GoogleCallback)
		}

		// Other resource routes can be added here
		// e.g., cycleRoutes, userRoutes, etc.
	}
}
