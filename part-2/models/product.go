package models

// ProductJSON model json file
type ProductJSON struct {
	ProductID string `json:"productId"`
	Image     string `json:"image"`
}

// Product model response json file
type Product struct {
	ProductID string   `json:"productId"`
	Images    []string `json:"images"`
}
