package main

import "fmt"

func calc(num1, num2 int) (func() int, func() int) {

	plus := func() int {
		return num1 + num2
	}

	minus := func() int {
		return num1 - num2
	}

	return plus, minus
}

func main() {

	plus, minus := calc(30, 10)

	fmt.Println(plus())  // 40
	fmt.Println(minus()) // 20
}
