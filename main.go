package main

import (
	"github.com/photo_catalog/internal/handlers"
)

func main() {
	router := handlers.SetupRouter()
	router.Run("localhost:8080")
}
