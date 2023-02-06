package main

import (
	"fmt"
)

func build(x func(int, int) int, a, b int) int {
	return x(a, b)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func main() {

	// fmt.Printf("%T\n", add)

	fmt.Println(build(add, 2, 4))
	fmt.Println(build(sub, 2, 4))

	// Anonym function
	// func(){}()

	// echo := func() {
	// 	fmt.Println("anonym")
	// }

	// echo()

	// square := func(x int) int {
	// 	return x * x
	// }

	// fmt.Println(square(10))
}
