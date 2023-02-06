package main

import (
	"fmt"
)

func calc(operation string, a, b float64) (bool, float64, string) {
	var err bool = false    // error handle
	var message string = "" // message for error
	var result float64      // result calculation

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			message = "Error: Can not divide to zero (0)"
			err = true
		} else {
			result = a / b
		}
	default:
		err = true
		message = "Please enter correct operation"
	}

	return err, result, message
}

func main() {
	err, result, message := calc("/", 10, 4)

	if err {
		fmt.Println(message)
	} else {
		fmt.Println(result) // 6
	}
}
