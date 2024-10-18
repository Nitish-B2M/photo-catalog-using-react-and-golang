package handlers

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/photo_catalog/internal/db"
	"github.com/photo_catalog/internal/handlers/catalogs"
	"github.com/photo_catalog/internal/handlers/health"
	"github.com/photo_catalog/internal/handlers/user"
	"github.com/photo_catalog/internal/queries"
	"github.com/photo_catalog/pkg/v1/services"
)

func SetupRouter() *gin.Engine {
	dbconn := db.ConnectToDB()
	dbAccess := queries.NewPersistentSQLDBStore(dbconn)

	router := gin.Default()
	// CORS configuration
	// corsConfig := cors.Config{
	// 	// AllowAllOrigins:  true,

	// 	AllowOrigins:     []string{"http://localhost:5173"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Application-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,

	// }

	// router.Use(cors.New(corsConfig))

	corsConfig := CORS()
	router.Use(corsConfig)
	router.Static("/assets", "D:/data/golang/learning/catalog/assets")

	router.GET("/health", health.HealthCheck)
	api := router.Group("/api/v1")

	catalogService := services.NewCatalogService(dbAccess)
	catalogHandler := catalogs.NewCatalogHandler(catalogService)
	v1CatalogGroup := api.Group("/catalog")
	{
		v1CatalogGroup.GET("/list", catalogHandler.ListCatalog)
		v1CatalogGroup.POST("/add", catalogHandler.AddCatalog)
	}

	userService := services.NewUserService(dbAccess)
	userHandler := user.NewUserHandler(userService)
	v1UserGroup := api.Group("/auth")
	{
		v1UserGroup.POST("/register", userHandler.Register)
		v1UserGroup.POST("/login", userHandler.Login)
		v1UserGroup.GET("/logout", userHandler.Logout)
		v1UserGroup.GET("/user", userHandler.User)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": []string{"404 page not found."}})
	})
	return router
}

func CORS() gin.HandlerFunc {
	config := cors.Config{}
	config.AllowHeaders = []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token", "Access-Control-Allow-Headers", "Access-Control-Request-Method", "Origin"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"}
	// config.AllowAllOrigins = true
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowBrowserExtensions = true
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.MaxAge = time.Hour * 12
	return cors.New(config)
}
