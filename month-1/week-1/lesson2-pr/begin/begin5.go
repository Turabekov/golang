package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 4

	v := math.Pow(float64(a), 3)
	s := 6 * math.Pow(float64(a), 6)

	fmt.Printf("V = %.0f S = %.0f", v, s)
}
