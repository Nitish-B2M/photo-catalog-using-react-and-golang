package queries

import (
	"context"
	"strings"

	"github.com/photo_catalog/internal/entities"
)

func (ms *persistentSQLDBStore) ListCatalogQuery(ctx context.Context, userID string) ([]entities.CatalogItem, error) {
	var items []entities.CatalogItem
	if err := ms.db.WithContext(ctx).Where("is_deleted = ? and publisher_id = ?", 0, userID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (ms *persistentSQLDBStore) AddCatalogQuery(ctx context.Context, addRequest *entities.CatalogItem) error {
	result := ms.db.WithContext(ctx).Create(&addRequest)
	err := result.Error
	if err != nil {
		if strings.Contains(err.Error(), "consts.ErrDuplicateEntry") {
			// _ := fmt.Sprintf("consts.ErrAccountAlreadyExists: %s", addRequest.Caption)
			return err
		} else {
			return err
		}
	}

	return nil
}
