package routes

import (
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	maxPhotoSize = 2 << 20 //2MB
)

var allowedFileTypes = map[string]struct{}{
	"image/jpeg": {},
	"image/png":  {},
	"image/jpg":  {},
	"image/gif":  {},
}

func UploadPhoto(c *gin.Context) (string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get photos"})
		return "", err
	}

	if err := ValidatePhoto(file); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return "", err
	}

	filename := strings.Split(file.Filename, ".") // need to implement extension validation
	filePath := "assets/uploads/" + filename[0] + "_" + time.Now().String() + filename[1]
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save."})
		return "", err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
	return filePath, nil
}

func ValidatePhoto(file *multipart.FileHeader) error {
	if file.Size > maxPhotoSize {
		return errors.New("file size exceeds the limit")
	}

	fileType := file.Header.Get("Content-Type")
	if _, ok := allowedFileTypes[fileType]; !ok {
		return errors.New("invalid file type; only JPEG, PNG, and GIF are allowed")
	}

	// implement in future
	// if err := checkImageDimensions(file); err != nil {
	// 	return err
	// }

	return nil
}

// func checkImageDimensions(file *multipart.FileHeader) error {
// 	src, err := file.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer src.Close()

// 	img, _, err := image.Decode(src)
// 	if err != nil {
// 		return err
// 	}

// 	// Check dimensions (e.g., maximum dimensions 4000x4000)
// 	if img.Bounds().Dx() > 4000 || img.Bounds().Dy() > 4000 {
// 		return errors.New("image dimensions exceed the limit of 4000x4000")
// 	}

// 	return nil
// }

func RemovePhoto(c *gin.Context) {
	filePath := c.Param("filepath")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	ok := os.Remove(filePath)
	log.Println(ok)

	c.JSON(http.StatusOK, gin.H{"message": "Image delete successfully"})
}
