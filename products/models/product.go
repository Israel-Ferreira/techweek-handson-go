package models

type Product struct {
	ID          int
	Sku         string
	Title       string
	Description string
}

func (p *Product) IsValid() error {
	return nil
}

func NewProduct() Product {
	return Product{}
}
