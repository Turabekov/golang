package main

import "fmt"

func main() {
	var a = 10
	var b = 5

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// a = a + b
	// b = a - b
	// a = a - b

	a, b = b, a

	fmt.Printf("After changing: a = %d b = %d\n", a, b)

}

// homework #1
// byte
// rune
