package entities

var _table_catalog = "catalog_items"

type CatalogItem struct {
	ID          string `json:"id" gorm:"primaryKey;type:varchar(36)" validate:"required,uuid"`
	Caption     string `json:"caption" gorm:"type:varchar(255)" validate:"required,min=1,max=255"`
	PublisherID string `json:"publisher_id" gorm:"type:varchar(36)" validate:"required,uuid"`
	ImagePath   string `json:"image" gorm:"type:varchar(255)" validate:"required,url"`
	Tags        string `json:"tags" gorm:"type:varchar(255)" validate:"max=255"`
	Location    string `json:"location" gorm:"type:varchar(100)" validate:"max=100"`
	IsDeleted   bool   `json:"is_deleted" gorm:"default:false"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func (m *CatalogItem) TableName() string {
	return _table_catalog
}
