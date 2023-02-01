package main

import (
	"fmt"
)

type son int
type complexFunc func(son, son) son

func direction() (complexFunc, complexFunc) {

	return func(a, b son) son {
			return a + b
		}, 
		func(a, b son) son {
			return a - b
		}
}

func main() {

	plus, minus := direction();

	fmt.Println(plus(20, 40))
	fmt.Println(minus(40, 50))
}
