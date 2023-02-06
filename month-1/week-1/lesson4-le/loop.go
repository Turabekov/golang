package main

import "fmt"

func main() {

	// Loop
	// 1. C-style loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 2. for only condition
	i := 0
	for i < 10 {

		fmt.Println(i)
		i++
	}

	// 3. Infinte loop
	for {
		fmt.Println("Infinite loop...")
	}

	// 4. Do-while loop
	// i := 0
	// for {

	// 	fmt.Println(i)
	// 	i++

	// 	if i > 10 {
	// 		break
	// 	}
	// }
}

// factorial(3)

// 3 = 1 * 2 * 3 = 6
