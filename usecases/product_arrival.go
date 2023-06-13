package usecases

import (
	"bufio"
	"fmt"
	"os"
	"project-go-fundamentals-example/models"
	"strconv"
	"strings"
)

//Explanation UseCase: ProductsArrival
//After I go to supermarket, I need to track all products I bought
//and where its are saved

//Process1: get product details from a text file
//Process2: Save products into a Slice within a storer (in that case, a foodStorage)
//Process3: List and show all saved products into the storer

//Also, a closure (or anonymous) function is implemented

func ProductsArrival(filename string, storer models.Storer) {
	products := getProductsFromFile(filename)
	storer.AddProducts(products)
	listProducts(storer)
}

func getProductsFromFile(filename string) []models.Product {
	var products []models.Product
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if line := fileScanner.Text(); line != "" { //if line is empty, go to next
			//fmt.Println("line: ", line)
			info := func() (string, float64) {
				//first line is name
				nameLine := strings.Split(line, "=")
				fileScanner.Scan() //read next line: volume
				volumeLine := strings.Split(fileScanner.Text(), "=")
				volume, _ := strconv.ParseFloat(volumeLine[1], 8)
				return nameLine[1], volume
			}
			products = append(products, models.NewProduct(info()))
		}
	}
	readFile.Close()
	return products
}

func listProducts(storer models.Storer) {
	if foodStorage, ok := storer.(*models.FoodStorage); ok {
		fmt.Println("The foodStorage has:")
		for _, product := range foodStorage.Products {
			fmt.Println(product.Name)
		}
	}
}
