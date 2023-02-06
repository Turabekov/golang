package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 5. Sig'imi 10ta bo'lgan slice e'lon qiling va unga random sonlar soling.
// Faqat toq indexda turgan elementlarni ekrangi chiqaruvchi function tuzing.

func getOddIndexes(arr []int) []int {
	newArr := []int{}

	for index, value := range arr {
		if index%2 != 0 {
			newArr = append(newArr, value)
		}
	}
	return newArr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(10)

	fmt.Println("Random array:", arr)
	fmt.Println("Odd index numbers:", getOddIndexes(arr))
}
