package mockproduct

import (
	"math/rand"

	productGenerator "github.com/NatthaphongS/go-basic/mock_product/internal/product_generator"
	"github.com/google/uuid"
)

// Product represents a product with name and price
type Product struct {
	PID string
	Name  string
	Price float64
}

// public function start with capital letter
func GetRandomProduct() Product {
	pid := uuid.New().String()
	// products := generateProductList()
	products := productGenerator.GenerateProductList() // use internal package
	
	return Product{
		PID: pid,
		Name:  products[rand.Intn(len(products))],
		Price: getRandomPrice(),
	}
}

// private function (can use only in this package) start with small letter
func generateProductList() []string{
	products := []string{
		"iPhone 15",
		"Samsung Galaxy S24",
		"MacBook Pro",
		"Dell XPS 13",
		"iPad Air",
		"Nintendo Switch",
		"PlayStation 5",
		"AirPods Pro",
		"Apple Watch",
		"Surface Pro",
	}

	return products
}