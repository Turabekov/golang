// 6. Sig'imi 10ta bo'lgan slice e'lon qiling va unga random sonlar soling.
// Shu slicening toq indexsidagi sonlarning yig'indisini bilan juft indexsidagi
// sonlar yig'indisini taqqoslab kattasini qaytaruvchi function tuzing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getMaxSumIndex(arr []int) (int, string) {
	sumEven := 0
	sumOdd := 0

	for index, value := range arr {
		if index%2 != 0 {
			sumOdd += value
		} else if index%2 == 0 {
			sumEven += value
		}
	}

	fmt.Println("sum even:", sumEven)
	fmt.Println("sum odd:", sumOdd)

	if sumEven > sumOdd {
		return sumEven, "even"
	} else {
		return sumOdd, "odd"
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(10)

	fmt.Println("Random array:", arr)
	num, msg := getMaxSumIndex(arr)
	fmt.Printf("Max sum %s indexes are greater: %d\n", msg, num)
}
