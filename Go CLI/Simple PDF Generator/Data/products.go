package data

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)


type Product struct {
	Name string 
	Description string
	Price float64
}
func populateProduct() Product {
	return Product{
		Name: gofakeit.ProductName(),
		Description: gofakeit.ProductDescription(),
		Price: gofakeit.Price(1, 50),
	}
}

func generateData() []string {
	f := populateProduct()
	froot := []string{}
	froot = append(froot, f.Name)
	froot = append(froot, f.Description)
	froot = append(froot, fmt.Sprintf("$%.2f", f.Price))
	return froot
}

func FruitList(len int) [][]string {
	var fruits [][]string
	for  i := 0; i < len; i++ {
		f := generateData()
		fruits = append(fruits, f)
	}
	return fruits
}