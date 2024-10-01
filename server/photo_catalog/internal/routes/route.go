package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	catalog := api.Group("/catalog")
	{
		// catalog.GET("", Getcatalogs)
		catalog.POST("", AddToCatalog)
	}
}
