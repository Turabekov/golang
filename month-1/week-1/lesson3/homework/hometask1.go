package main

import "fmt"

func direction(num1, num2 int) (int, int) {
	return num1 * 2, num2 * 2
}

func main() {
	L, R := direction(10, 15)

	fmt.Println(L) // 20
	fmt.Println(R) // 30

}
