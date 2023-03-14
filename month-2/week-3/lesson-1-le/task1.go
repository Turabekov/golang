package main

import "fmt"

func Add(a, b int, addChan chan int) {
	addChan <- a + b
}

func Sub(a, b int, subChan chan int) {
	subChan <- a - b
}

func Mult(a, b int, multChan chan int) {
	multChan <- a * b
}

func Div(a, b int, divChan chan float64) {
	if b == 0 {
		divChan <- 0
	} else {
		divChan <- float64(a) / float64(b)
	}
}

func main() {
	var addChan = make(chan int)
	var subChan = make(chan int)
	var multChan = make(chan int)
	var divChan = make(chan float64)

	a, b := 10, 5

	go Add(a, b, addChan)
	go Sub(a, b, subChan)
	go Mult(a, b, multChan)
	go Div(a, b, divChan)

	fmt.Println(<-addChan)
	fmt.Println(<-subChan)
	fmt.Println(<-multChan)
	fmt.Println(<-divChan)
}
