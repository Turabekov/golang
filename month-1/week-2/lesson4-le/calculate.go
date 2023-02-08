package main

import (
	"fmt"
)

type Operation uint8

const (
	Additional Operation = iota + 132
	Sub
	Div
	Mult
)

func calc(a, b int, c Operation) int {

	if c == Additional {
		return a + b
	} else if c == Sub {
		return a - b
	} else if c == Div {
		return a / b
	} else if c == Mult {
		return a * b
	}

	return 0
}

func main() {

	fmt.Println("Add: ", calc(3, 4, Additional))
	fmt.Println("Sub: ", calc(3, 4, Sub))
	fmt.Println("Div: ", calc(3, 4, Div))
	fmt.Println("Mult: ", calc(3, 4, Mult))
}
