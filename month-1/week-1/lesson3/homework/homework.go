package main

import "fmt"

func main () {

	plus, minus := calc(30, 10)

	fmt.Println(plus())   // 40
	fmt.Println(minus())  // 20
}
