package main

import "fmt"

var global = "global variable"

// niladic
func welcome() {

	fmt.Println("Welcome")
}

// with return funciton
func isActive() int {

	return 1
}

// argument function
func add(firstNumber, secondNumber float64) float64 {

	return firstNumber + secondNumber
}

func plus(x int, y float64) float64 {

	// type cast
	return float64(x) + y
}

// multiple return
func Color() (uint8, uint8, uint8) {

	return 100, 255, 108
}

func main() {

	// var a, b = 1, 4
	// a, b, z := 10, 12, 14

	// fmt.Println(a, b, z)


	R, _, B := Color()

	fmt.Println(R, B)
}
