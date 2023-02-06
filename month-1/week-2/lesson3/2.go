package main

import (
	"fmt"
	"strconv"
)

func add(a, b int) int  { return a + b }
func sub(a, b int) int  { return a - b }
func div(a, b int) int  { return a / b }
func mult(a, b int) int { return a * b }
func pow(a, b int) int {
	p := a
	for i := 1; i < b; i++ {
		a = a * p
	}

	return a
}
func prosent(a, b int) int { return int((float64(a) / 100) * float64(b)) }

func main() {
	var calcMap = map[string]func(int, int) int{
		"+":  add,
		"-":  sub,
		"/":  div,
		"*":  mult,
		"**": pow,
		"%":  prosent,
	}

	var expressions = [][]string{
		{"1", "+", "3"},
		{"4", "-", "10"},
		{"10", "/", "2"},
		{"2", "*", "12"},
		{"one", "+", "two"},
		{"4", "div", "10"},
		{"123456", "+", "98765"},
		{"5", "**", "2"},
		{"20", "%", "10"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression")
			continue
		}

		first_number, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		operation := expression[1]

		onFunc, ok := calcMap[operation]

		if !ok {
			fmt.Println("unsupported operation")
			continue
		}
		second_number, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := onFunc(first_number, second_number)
		fmt.Println(result)
	}

}
