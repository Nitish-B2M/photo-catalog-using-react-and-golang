package queries

import (
	"context"

	"github.com/photo_catalog/internal/entities"
	"github.com/photo_catalog/pkg/v1/requests"
	"github.com/photo_catalog/pkg/v1/responses"
)

type MockPersistentSQLDBStore struct{}

var _ PersistentSQLDBStorer = (*MockPersistentSQLDBStore)(nil)

func (ms *MockPersistentSQLDBStore) AddCatalogQuery(ctx context.Context, addRequest *entities.CatalogItem) error {
	return nil
}

func (ms *MockPersistentSQLDBStore) ListCatalogQuery(ctx context.Context, userID string) ([]entities.CatalogItem, error) {
	return nil, nil
}

// authentication logic
func (ms *MockPersistentSQLDBStore) RegisterQuery(ctx context.Context, addUser *entities.UserRegister) error {
	return nil
}

func (ms *MockPersistentSQLDBStore) LoginQuery(ctx context.Context, addUser *requests.UserLoginRequest) (responses.LoginUserResponse, error) {
	var lur responses.LoginUserResponse
	return lur, nil
}

func (ms *MockPersistentSQLDBStore) LogoutQuery(ctx context.Context, userID string) (string, error) {
	return "", nil
}

func (ms *MockPersistentSQLDBStore) UserQuery(ctx context.Context, userID string) (responses.LoggedInUResponse, error) {
	var user responses.LoggedInUResponse
	return user, nil
}

func (ms *MockPersistentSQLDBStore) UserActivityQuery(ctx context.Context, userID string) (responses.LoggedInUResponse, error) {
	var user responses.LoggedInUResponse
	return user, nil
}
