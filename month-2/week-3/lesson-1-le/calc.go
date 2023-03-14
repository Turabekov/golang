package main

import "fmt"

func Square(input int, num chan int) {
	num <- input * input
}

func Cube(input int, num chan int) {
	num <- input * input * input
}

func main() {
	var squareChan = make(chan int)
	var cubeChan = make(chan int)
	var input = 20

	go Square(input, squareChan)
	go Cube(input, cubeChan)

	fmt.Println(<-squareChan)
	fmt.Println(<-cubeChan)
}
