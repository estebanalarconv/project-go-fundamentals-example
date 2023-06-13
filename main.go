package main

import (
	"project-go-fundamentals-example/models"
	"project-go-fundamentals-example/usecases"
)

func main() {
	//creating a storer: foodStorage, with capacity 30.000 ml, or 30 liter
	foodStorage := models.NewFoodStorage(30000)
	usecases.ProductsArrival("usecases/product_examples.txt", &foodStorage)
}
