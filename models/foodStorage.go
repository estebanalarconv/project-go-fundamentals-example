package models

type FoodStorage struct {
	Products []Product
	Capacity float64
}

func NewFoodStorage(capacity float64) FoodStorage {
	return FoodStorage{
		Products: make([]Product, 0, 10),
		Capacity: capacity,
	}
}

func (f FoodStorage) CalculateFreeCapacity() float64 {
	total := 0.0
	for _, product := range f.Products {
		total += product.Volume
	}
	return f.Capacity - total
}

func (f *FoodStorage) AddProducts(newProducts []Product) {
	f.Products = append(f.Products, newProducts...)
}
