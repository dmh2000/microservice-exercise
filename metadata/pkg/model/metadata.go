package model

// movie metadata
type Metadata struct {
	ID          string `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
}
