package models

type Product struct {
	Name   string
	Volume float64 //ml
}

func NewProduct(name string, volume float64) Product {
	return Product{
		Name:   name,
		Volume: volume,
	}
}

func (p Product) IsExpired() bool {
	return false
}
