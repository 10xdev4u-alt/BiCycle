package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/api"
	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/config"
	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/database"
	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/routes"
)

func main() {
	cfg := config.LoadConfig()

	db := database.InitDB(cfg)
	defer db.Close()

	api := api.NewAPI(db, cfg)

	router := gin.Default()

	// Setup routes
	routes.SetupRouter(router, api)

	// Add a root endpoint for basic health check
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "The BiCycle App API is running!",
		})
	})

	router.Run(fmt.Sprintf(":%s", cfg.AppPort)) // listen and serve on configured port
}