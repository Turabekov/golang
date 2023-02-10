package main

import "fmt"

func main() {

	var abstractVariable interface{} = 2134253

	// Type Switch
	switch abstractVariable.(type) {
	case int:
		// Type Assertion
		fmt.Println("Int:", abstractVariable.(int))
	case string:
		fmt.Println("String:", abstractVariable.(string))
	default:
		fmt.Println("Unsupported type")
	}

	// Comma-ok
	if _, ok := abstractVariable.(int); ok {
		fmt.Println("Int:", abstractVariable.(int))
	}
}
