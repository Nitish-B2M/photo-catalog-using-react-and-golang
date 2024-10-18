package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/photo_catalog/pkg/v1/services"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	User(c *gin.Context)
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{service: service}
}

func (srv *userHandler) Register(c *gin.Context) {
	resp, err := srv.service.Register(c)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (srv *userHandler) Login(c *gin.Context) {
	resp, err := srv.service.Login(c)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusFound, resp)
}

func (srv *userHandler) User(c *gin.Context) {
	resp, err := srv.service.User(c)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusFound, resp)
}

func (srv *userHandler) Logout(c *gin.Context) {
	resp, err := srv.service.Logout(c)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
