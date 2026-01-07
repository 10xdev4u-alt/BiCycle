package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/config"
)

func main() {
	cfg := config.LoadConfig()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, PrinceTheProgrammer! Welcome to TheBiCycleApp Backend (with Config Loading)!",
		})
	})

	router.Run(fmt.Sprintf(":%s", cfg.AppPort)) // listen and serve on configured port
}