package catalogs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/photo_catalog/pkg/v1/services"
)

type CatalogHandler interface {
	AddCatalog(c *gin.Context)
	ListCatalog(c *gin.Context)
}

type catalogHandler struct {
	service services.CatalogService
}

func NewCatalogHandler(service services.CatalogService) CatalogHandler {
	return &catalogHandler{service: service}
}

func (srv *catalogHandler) AddCatalog(c *gin.Context) {
	resp, err := srv.service.AddCatalog(c)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (srv *catalogHandler) ListCatalog(c *gin.Context) {
	resp, err := srv.service.ListCatalog(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusFound, resp)
}
