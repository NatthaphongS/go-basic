package productgenerator

import (
	"math/rand"
)

// GenerateRandomPrice generates a random price between min and max
func GenerateRandomPrice(min, max int) float64 {
	return float64(rand.Intn(max-min+1) + min)
}

// GenerateProductList returns a list of sample products
func GenerateProductList() []string {
	return []string{
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
		"Google Pixel 8",
		"Microsoft Surface Laptop",
		"Sony WH-1000XM5",
		"Tesla Model S",
		"Canon EOS R5",
	}
}