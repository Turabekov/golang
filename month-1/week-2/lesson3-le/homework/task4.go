package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 4. Sig'imi 10ta bo'lgan slice e'lon qiling va unga random sonlar soling.
// Shu sonlarni teskari tartibda chiqaring.

func reverseSlice(arr []int) []int {
	reverseArr := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		reverseArr = append(reverseArr, arr[i])
	}

	return reverseArr

}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(10)

	fmt.Println("Random array:", arr)
	fmt.Println("Reverse numbers:", reverseSlice(arr))
}
