package main

import (
	"fmt"
	"math"
)

func main() {
	var A int = 2

	fmt.Println("A to the power of 2 =", int(math.Pow(float64(A), 2)))
	fmt.Println("A to the power of 3 =", int(math.Pow(float64(A), 3)))
	fmt.Println("A to the power of 5 =", int(math.Pow(float64(A), 5)))
	fmt.Println("A to the power of 10 =", int(math.Pow(float64(A), 10)))
	fmt.Println("A to the power of 15 =", int(math.Pow(float64(A), 15)))
}
