package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/photo_catalog/internal/entities"
)

const maxPhotoSize = 5 * 1024 * 1024

func SaveFileToLocal(c *gin.Context, file *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(file.Filename)
	timestamp := time.Now().Format("20060102150405")
	filename := strings.TrimSuffix(file.Filename, ext) + "_" + timestamp + ext
	filePath := filepath.Join("assets/uploads/", filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}

// isValidExtension checks if the file extension is in the allowed list.
func isValidExtension(ext string, allowed []string) bool {
	for _, allowedExt := range allowed {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

func ValidatePhoto(file *multipart.FileHeader) error {
	if file == nil {
		return errors.New("file is nil")
	}

	if file.Size > maxPhotoSize {
		return errors.New("file size exceeds the limit")
	}

	ext := filepath.Ext(file.Filename)
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp"}
	if !isValidExtension(ext, allowedExtensions) {
		return errors.New(fmt.Sprintf("invalid file extension: %s", ext))
	}

	return nil
}

func SendFullPathTOResponse(data []entities.CatalogItem) []entities.CatalogItem {
	// basePath := "D:/data/golang/learning/catalog/"
	basePath := ""

	for i := range data {
		relativePath := data[i].ImagePath
		if strings.TrimSpace(relativePath) == "" || filepath.Ext(relativePath) == "" {
			continue
		}

		fullPath := filepath.Join(basePath, relativePath)
		data[i].ImagePath = fullPath
	}
	return data
}
