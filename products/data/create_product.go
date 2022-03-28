package data

type CreateProduct struct {
	Title       string `json:"title"`
	Sku         string `json:"sku"`
	Description string `json:"description"`
	Brand       string `json:"brand"`
}
