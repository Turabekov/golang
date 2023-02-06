package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 3. Sig'imi 10ta bo'lgan slice e'lon qiling va unga random sonlar soling.
// Shu sonlar ichidan 2-tub sonni qaytaruvchi funksiya tuzing.

func getSecondPrimeNum(arr []int) int {
	storage := []int{}
	count := 0

	for i := 0; i < len(arr); i++ {
		count = 0
		for j := 1; j <= arr[i]; j++ {
			if arr[i]%j == 0 {
				count++
			}
		}
		if count == 2 {
			storage = append(storage, arr[i])
		}
	}

	return storage[1]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(10)

	fmt.Println("Random array:", arr)
	fmt.Println("Second prime number:", getSecondPrimeNum(arr))
}
