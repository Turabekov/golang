package main

import (
	"fmt"
)

func swap(x, y *int) {
	*x, *y = *y, *x

	fmt.Println("x:", x)
	fmt.Println(*x, "::::")
}

func Add(a, b, result *int) {
	*result = *a + *b
}

// Sub()
// Div()
// Mult()

func main() {

	var a, b, result = 32, 43, 0
	Add(&a, &b, &result)
	fmt.Println(result)

	// var a, b = 10, 15

	// fmt.Println("a:", &a)

	// swap(&a, &b)

	// fmt.Println(a, b)
}
