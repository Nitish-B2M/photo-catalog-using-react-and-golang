package services

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/photo_catalog/internal/entities"
	"github.com/photo_catalog/internal/queries"
	"github.com/photo_catalog/pkg/utils"
	"github.com/photo_catalog/pkg/v1/requests"
	"github.com/photo_catalog/pkg/v1/responses"
)

type catalogService struct {
	db queries.PersistentSQLDBStorer
}

type CatalogService interface {
	AddCatalog(c *gin.Context) (responses.Response, responses.ErrorMessage)
	ListCatalog(c *gin.Context) (responses.Response, error)
}

func NewCatalogService(dbacess queries.PersistentSQLDBStorer) CatalogService {
	return &catalogService{db: dbacess}
}

// implementing method
func (s catalogService) AddCatalog(c *gin.Context) (responses.Response, responses.ErrorMessage) {
	var responseData responses.Response
	var errorResponse responses.ErrorMessage

	caption := c.PostForm("caption")
	location := c.PostForm("location")
	tags := c.PostForm("tags")

	userID, err := utils.GetContentFromCookie(c, "token")
	if err != nil {
		errorResponse := utils.NewErrorMessage("Unauthenticated", "Unauthenticated user", err.Error())
		return responseData, errorResponse
	}

	addCatalogRequest := requests.CatalogAddRequest{
		RequestCaption:     caption,
		RequestLocation:    location,
		RequestTags:        tags,
		RequestPublisherID: userID,
	}

	validationErrors := utils.ValidateStruct(addCatalogRequest)
	if len(validationErrors) > 0 {
		errorResponse := utils.NewErrorMessage("ValidationError", "Input validation failed", validationErrors)
		return responseData, errorResponse
	}

	file, err := c.FormFile("image")
	if err != nil {
		errorResponse := utils.NewErrorMessage("UploadingError", "Not able to fetch file", err)
		return responseData, errorResponse
	}

	if err := utils.ValidatePhoto(file); err != nil {
		errorResponse := utils.NewErrorMessage("UploadingError", "File validation failed", err)
		return responseData, errorResponse
	}

	filePath, err := utils.SaveFileToLocal(c, file)
	if err != nil {
		errorResponse := utils.NewErrorMessage("StoringError", "File storing failed", err)
		return responseData, errorResponse
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	addCatalogItem := entities.CatalogItem{
		ID:          uuid.New().String(),
		Caption:     addCatalogRequest.RequestCaption,
		PublisherID: addCatalogRequest.RequestPublisherID,
		ImagePath:   filePath,
		Tags:        addCatalogRequest.RequestTags,
		Location:    addCatalogRequest.RequestLocation,
	}

	err = s.db.AddCatalogQuery(ctx, &addCatalogItem)
	if err != nil {
		errorResponse := utils.NewErrorMessage("StoringDBError", "File storing failed", err)
		return responseData, errorResponse
	}

	var newCatalogItem []responses.CatalogItemResponse
	newCatalog := append(newCatalogItem, responses.CatalogItemResponse{
		Caption:     addCatalogRequest.RequestCaption,
		PublisherID: addCatalogRequest.RequestPublisherID,
		ImagePath:   filePath,
		Tags:        addCatalogRequest.RequestTags,
		Location:    addCatalogRequest.RequestLocation,
	})

	responseData.Data = newCatalog
	responseData.Message = fmt.Sprintf("Item '%s' added to catalog successfully", addCatalogItem.Caption)
	responseData.RecordSet = nil
	return responseData, errorResponse
}

func (s catalogService) ListCatalog(c *gin.Context) (responses.Response, error) {
	var responseData responses.Response

	userID, err := utils.GetContentFromCookie(c, "token")
	if err != nil {
		_ = utils.NewErrorMessage("Unauthenticated", "Unauthenticated user", err.Error())
		return responseData, err
	}

	responseData.Message = "Memories"
	item, err := s.db.ListCatalogQuery(c, userID)
	if err != nil {
		return responseData, err
	}

	item = utils.SendFullPathTOResponse(item)

	responseData.Data = item
	responseData.RecordSet = ""
	return responseData, nil
}
