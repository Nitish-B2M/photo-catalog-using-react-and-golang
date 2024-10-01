package response

// CatalogItem represents the structure of the item to be added to the catalog
type CatalogItem struct {
	ID        string `json:"id"`
	Caption   string `json:"caption"`
	Location  string `json:"location"`
	ImagePath string `json:"image"`
}
