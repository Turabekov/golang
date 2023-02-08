package main

import "fmt"

const PI float64 = 3.14
const (
	A = iota * 2
	B
	C
	D
	E
	F
)

func main() {

	const R = 8.31

	fmt.Println(PI * 3)

	fmt.Println(A, B, C, D, E, F)
	fmt.Printf("%T\n", A)
}
