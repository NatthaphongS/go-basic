package main

import (
	"fmt"

	mockProduct "github.com/NatthaphongS/go-basic/mock_product"
)

func main() {
	// fmt.Println("Hello, World!")
	product := mockProduct.GetRandomProduct()
	fmt.Println(product.Price)
	fmt.Printf("PID: %s\nProductName: %s\nPrice: %f\n",product.PID, product.Name, product.Price)
}