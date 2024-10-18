package queries

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/photo_catalog/internal/entities"
	"github.com/photo_catalog/pkg/utils"
	"github.com/photo_catalog/pkg/v1/requests"
	"github.com/photo_catalog/pkg/v1/responses"
	"gorm.io/gorm"
)

func (ms *persistentSQLDBStore) RegisterQuery(ctx context.Context, addUser *entities.UserRegister) error {
	result := ms.db.WithContext(ctx).Create(&addUser)
	err := result.Error
	if err != nil {
		if err, ok := err.(*mysql.MySQLError); ok && err.Number == 1062 {
			return errors.New("account already exists")
		}
		return err
	}

	return nil
}

// check user exist
func CheckUserExistsByEmail(ms *persistentSQLDBStore, ctx context.Context, email string) (entities.UserRegister, error) {
	var user entities.UserRegister
	err := ms.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user does not exist, invalid request")
		}
		return user, err
	}
	return user, nil
}

func (ms *persistentSQLDBStore) LoginQuery(ctx context.Context, loginUser *requests.UserLoginRequest) (responses.LoginUserResponse, error) {

	var userResponse responses.LoginUserResponse
	user, err := CheckUserExistsByEmail(ms, ctx, loginUser.RequestEmail)
	if err != nil {
		return userResponse, err
	}

	if !utils.VerifyData(user.Password, []byte(loginUser.RequestPassword)) {
		return userResponse, errors.New("invalid password")
	}

	userResponse.ID = user.ID
	userResponse.Email = user.Email
	userResponse.Username = user.Username

	return userResponse, nil
}

func (ms *persistentSQLDBStore) LogoutQuery(ctx context.Context, userID string) (string, error) {
	if err := ms.db.WithContext(ctx).Table("users").Where("id = ?", userID).Update("last_active", time.Now().Unix()).Error; err != nil {
		return "", err
	}

	if ms.db.RowsAffected == 0 {
		return "", fmt.Errorf("user ID %s not found", userID)
	}

	return "logout succcessfully", nil
}

func (ms *persistentSQLDBStore) UserQuery(ctx context.Context, userID string) (responses.LoggedInUResponse, error) {
	var user responses.LoggedInUResponse

	if err := ms.db.WithContext(ctx).Table("users").Where("id = ?", userID).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ms *persistentSQLDBStore) UserActivityQuery(ctx context.Context, userID string) (responses.LoggedInUResponse, error) {
	var user responses.LoggedInUResponse

	if err := ms.db.WithContext(ctx).Table("users").Where("id = ?", userID).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
