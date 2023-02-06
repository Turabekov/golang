package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64 = 1, 2, 4
	D := B*B - 4*A*C

	if D > 0 && A != 0 {
		x1 := (-B + math.Sqrt(D)) / (2 * A)
		x2 := (-B - math.Sqrt(D)) / (2 * A)

		fmt.Println("X1 =", x1, "X2 =", x2)
	} else if D == 0 {
		x := (-B + math.Sqrt(D)) / (2 * A)
		fmt.Println("X =", x)
	} else {
		fmt.Println("Please enter correct koeffitients")
	}
}
