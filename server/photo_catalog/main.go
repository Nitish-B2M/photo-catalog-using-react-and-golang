package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/photo_catalog/internal/routes"
)

func main() {
	router := gin.Default()

	// CORS configuration
	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour, // Correctly specifying MaxAge
	}

	router.Use(cors.New(corsConfig))
	routes.SetupRoutes(router)
	router.Run("localhost:8080")
}
