package modals

import (
	"time"
)

type CatalogItem struct {
	ID          string    `json:"id"`
	Caption     string    `json:"caption"`
	PublisherID string    `json:"publisher_id"`
	ImagePath   string    `json:"image"`
	Tags        string    `json:"tags"`
	Location    string    `json:"location"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"created_at" `
	UpdatedAt   time.Time `json:"updated_at"`
}
