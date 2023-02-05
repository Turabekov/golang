package main

import "fmt"

func main() {

	var v, u float64 = 3, 1
	var t1, t2 float64 = 1, 2

	if v > u {
		s1 := (v + u) / t1
		s2 := (v - u) / t2

		fmt.Println("S1 =", s1, "S2 =", s2)
	} else {
		fmt.Println("Error V should be greater than U")
	}
}
