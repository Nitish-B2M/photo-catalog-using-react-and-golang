package queries

import (
	"context"

	"github.com/photo_catalog/internal/entities"
	"github.com/photo_catalog/pkg/v1/requests"
	"github.com/photo_catalog/pkg/v1/responses"
	"gorm.io/gorm"
)

type persistentSQLDBStore struct {
	db *gorm.DB
}

type PersistentSQLDBStorer interface {
	// catalog
	ListCatalogQuery(ctx context.Context, userID string) ([]entities.CatalogItem, error)
	AddCatalogQuery(ctx context.Context, addRequest *entities.CatalogItem) error

	// authentication
	RegisterQuery(ctx context.Context, addUser *entities.UserRegister) error
	LoginQuery(ctx context.Context, addUser *requests.UserLoginRequest) (responses.LoginUserResponse, error)
	LogoutQuery(ctx context.Context, userID string) (string, error)
	UserQuery(ctx context.Context, userID string) (responses.LoggedInUResponse, error)
	UserActivityQuery(ctx context.Context, userID string) (responses.LoggedInUResponse, error)
}

func NewPersistentSQLDBStore(dbconn *gorm.DB) PersistentSQLDBStorer {
	return &persistentSQLDBStore{db: dbconn}
}
