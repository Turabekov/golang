package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 2. Sig'imi 10ta bo'lgan slice e'lon qiling va unga random sonlar soling.
// Shu sonlar yigindisini qaytaruvchi funksiya tuzing.

func sumElems(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	arr := make([]int, 10)
	// arr := rand.Perm(10)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(len(arr))
	}

	fmt.Println("Random array:", arr)
	fmt.Println("Summa:", sumElems(arr))
}
