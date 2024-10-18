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

type userService struct {
	db queries.PersistentSQLDBStorer
}

type UserService interface {
	Register(c *gin.Context) (responses.Response, responses.ErrorMessage)
	Login(c *gin.Context) (responses.Response, responses.ErrorMessage)
	Logout(c *gin.Context) (responses.Response, responses.ErrorMessage)
	User(c *gin.Context) (responses.Response, responses.ErrorMessage)
	UserActivity(c *gin.Context) (responses.Response, responses.ErrorMessage)
}

func NewUserService(dbacess queries.PersistentSQLDBStorer) UserService {
	return &userService{db: dbacess}
}

func (us userService) Register(c *gin.Context) (responses.Response, responses.ErrorMessage) {
	var responseData responses.Response
	var errorResponse responses.ErrorMessage
	var userRequest requests.UserRequest

	if err := c.BindJSON(&userRequest); err != nil {
		errorResponse := utils.NewErrorMessage("InvalidInput", "Request data is invalid", err.Error())
		return responseData, errorResponse
	}

	validationErrors := utils.ValidateStruct(userRequest)
	if len(validationErrors) > 0 {
		errorResponse := utils.NewErrorMessage("ValidationError", "Input validation failed", validationErrors)
		return responseData, errorResponse
	}

	if !utils.IsValidEmail(userRequest.RequestEmail) {
		errorResponse := utils.NewErrorMessage("InvalidInput", "Please input valid email", []string{"Invalid email id"})
		return responseData, errorResponse
	}

	hashedPassword, err := utils.HashData([]byte(userRequest.RequestPassword))
	if err != nil {
		errorResponse := utils.NewErrorMessage("HashingFailed", "Could not hash password", err.Error())
		return responseData, errorResponse
	}

	addUser := entities.UserRegister{
		ID:       uuid.New().String(),
		Username: userRequest.RequestUsername,
		Email:    userRequest.RequestEmail,
		Password: hashedPassword,
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err = us.db.RegisterQuery(ctx, &addUser)
	if err != nil {
		errorDesc := "An unexpected error occurred during registration."
		errorResponse := utils.NewErrorMessage("RegistrationError", errorDesc, err.Error())
		return responseData, errorResponse
	}

	var newUserResponse []responses.RegisterUserResponse
	newUser := append(newUserResponse, responses.RegisterUserResponse{
		ID:       uuid.New().String(),
		Username: userRequest.RequestUsername,
		Email:    userRequest.RequestEmail,
	})

	responseData.Data = newUser
	responseData.Message = fmt.Sprintf("User '%s' register successfully", addUser.Username)
	responseData.RecordSet = nil
	return responseData, errorResponse
}

func (us userService) Login(c *gin.Context) (responses.Response, responses.ErrorMessage) {
	var responseData responses.Response
	var errorData responses.ErrorMessage
	var loginRequest requests.UserLoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		errorResponse := utils.NewErrorMessage("InvalidInput", "Request data is invalid", err.Error())
		return responseData, errorResponse
	}

	if !utils.IsValidEmail(loginRequest.RequestEmail) {
		errorResponse := utils.NewErrorMessage("InvalidInput", "Please input valid email", []string{"Invalid email id"})
		return responseData, errorResponse
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userResponse, err := us.db.LoginQuery(ctx, &loginRequest)
	if err != nil {
		errorResponse := utils.NewErrorMessage("LoginError", "An unexpected error occurred during login", err.Error())
		return responseData, errorResponse
	}

	token, err := utils.GenerateTokenUsingClaims(userResponse.ID)
	if err != nil {
		errorResponse := utils.NewErrorMessage("LoginError", "An unexpected error occurred during login", err)
		return responseData, errorResponse
	}

	utils.SetCookie(c, utils.CookieConfig{Name: "token", Value: token, HttpOnly: true, MaxAge: int((24 * time.Hour).Seconds()), Secure: false})

	responseData.Data = userResponse
	responseData.Message = "User logged in"
	responseData.RecordSet = nil
	return responseData, errorData
}

func (us userService) Logout(c *gin.Context) (responses.Response, responses.ErrorMessage) {
	var responseData responses.Response
	var errorData responses.ErrorMessage

	userID, err := utils.GetContentFromCookie(c, "token")
	if err != nil {
		errorResponse := utils.NewErrorMessage("Unauthenticated", "Unauthenticated user", err.Error())
		return responseData, errorResponse
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	msg, err := us.db.LogoutQuery(ctx, userID)
	if err != nil {
		errorResponse := utils.NewErrorMessage("Unauthenticated", "Unauthenticated user", err.Error())
		return responseData, errorResponse
	}
	utils.SetCookie(c, utils.CookieConfig{Name: "token", Value: "", HttpOnly: true, MaxAge: int((-1 * time.Hour).Seconds()), Secure: false})

	responseData.Message = msg
	return responseData, errorData
}

func (us userService) User(c *gin.Context) (responses.Response, responses.ErrorMessage) {
	var responseData responses.Response
	var errorData responses.ErrorMessage

	userID, err := utils.GetContentFromCookie(c, "token")
	if err != nil {
		errorResponse := utils.NewErrorMessage("Unauthenticated", "Unauthenticated user", err.Error())
		return responseData, errorResponse
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userResponse, err := us.db.UserQuery(ctx, userID)
	if err != nil {
		errorResponse := utils.NewErrorMessage("Unauthenticated", "Unauthenticated user", err.Error())
		return responseData, errorResponse
	}

	responseData.Data = userResponse
	responseData.Message = "profile details"
	return responseData, errorData
}

func (us userService) UserActivity(c *gin.Context) (responses.Response, responses.ErrorMessage) {
	var responseData responses.Response
	var errorData responses.ErrorMessage

	userID, err := utils.GetContentFromCookie(c, "token")
	if err != nil {
		errorResponse := utils.NewErrorMessage("Unauthenticated", "Unauthenticated user", err.Error())
		return responseData, errorResponse
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userResponse, err := us.db.UserQuery(ctx, userID)
	if err != nil {
		errorResponse := utils.NewErrorMessage("Unauthenticated", "Unauthenticated user", err.Error())
		return responseData, errorResponse
	}

	responseData.Data = userResponse
	responseData.Message = "profile details"
	return responseData, errorData
}
