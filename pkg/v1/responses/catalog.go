package responses

type CatalogItemResponse struct {
	ID          string `json:"id"`
	Caption     string `json:"caption"`
	PublisherID string `json:"publisher_id"`
	ImagePath   string `json:"image"`
	Tags        string `json:"tags"`
	Location    string `json:"location"`
	IsDeleted   bool   `json:"is_deleted"`
	CreatedAt   int64  `json:"created_at" `
	UpdatedAt   int64  `json:"updated_at"`
}

var _table_catalog_item = "catalog_item"

// TableName get sql table name catalog_item
func (m *CatalogItemResponse) TableName() string {
	return _table_catalog_item
}
