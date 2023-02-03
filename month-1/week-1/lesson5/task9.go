package main

import (
	"fmt"
	"math"
)

// 9. N soni berilgan. Shu songacha bo’lgan 2 ning barcha darajalarini o’sish tartibida chiqaring.

func main() {
	var n int = 44

	for i := 1; ; i++ {
		num := math.Pow(2, float64(i))

		if num < float64(n) {
			fmt.Printf("%0.f ", math.Pow(2, float64(i)))
		} else {
			break
		}

	}
}
