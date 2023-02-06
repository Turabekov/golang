package main

import (
	"fmt"
	"math"
)

func main() {
	var l int = 2
	const pi = 3.14

	r := float64(l) / (2 * pi)
	s := pi * math.Pow(r, 2)

	fmt.Printf("Radius = %0.4f S = %0.4f", r, s)

}
