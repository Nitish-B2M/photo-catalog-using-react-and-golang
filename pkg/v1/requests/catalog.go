package requests

type CatalogAddRequest struct {
	RequestCaption     string `json:"caption"`
	RequestLocation    string `json:"location"`
	RequestImage       string `json:"image"`
	RequestTags        string `json:"tags"`
	RequestPublisherID string `json:"publisher_id"`
}
