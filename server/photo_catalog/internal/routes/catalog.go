package routes

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/photo_catalog/internal/db"
	"github.com/photo_catalog/internal/modals"
)

func AddToCatalog(c *gin.Context) {
	var item modals.CatalogItem
	item.ID = uuid.New().String()

	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if item.ImagePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image not provided"})
		return
	}

	if err := ValidateCatalogItem(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imgData, err := decodeBase64Image(item.ImagePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode image"})
		return
	}

	filePath := fmt.Sprintf("assets/uploads/%s_%d.jpg", item.ID[:8], time.Now().Unix())
	if err := os.MkdirAll("assets/uploads", os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create uploads directory"})
		return
	}

	if err := os.WriteFile(filePath, imgData, 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	item.ImagePath = filePath
	item.IsDeleted = false
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	if err := saveCatalogToDB(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save item to the database: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func decodeBase64Image(data string) ([]byte, error) {
	const prefix = "data:image/jpg;base64,"
	if len(data) > len(prefix) && data[:len(prefix)] == prefix {
		data = data[len(prefix):]
	}

	return base64.StdEncoding.DecodeString(data)
}

func ValidateCatalogItem(item *modals.CatalogItem) error {
	if item.PublisherID == "" {
		item.PublisherID = ""
	}
	if item.Caption == "" {
		item.Caption = ""
	}
	return nil
}

func getCatalogById(id string) (*modals.CatalogItem, bool) {
	return nil, false
}

func saveCatalogToDB(item *modals.CatalogItem) error {
	db, err := db.MysqlConnection()
	if err != nil {
		return err
	}

	// Save the item to the database
	if err := db.Create(item).Error; err != nil {
		return err
	}

	return nil
}

// ======================== Update Logic
func RemoveFromCatalog(c *gin.Context) {
	itemId := c.Param("id")
	catalog, isExist := getCatalogById(itemId)
	if !isExist {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not exist."})
		return
	}
	c.IndentedJSON(http.StatusOK, catalog)
}
