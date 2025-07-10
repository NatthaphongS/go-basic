package mockproduct

import "math/rand"

func getRandomPrice() float64 {
	min := 100
	max := 100000

	return float64(rand.Intn(max-min+1) + min)
}