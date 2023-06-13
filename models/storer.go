package models

type Storer interface {
	CalculateFreeCapacity() float64
	AddProducts([]Product)
}
