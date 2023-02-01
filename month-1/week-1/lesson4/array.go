package main

import "fmt"

func main() {

	// array vs slice
	// var nums = [10]int{1, 2, 4, 5, 6, 7}

	// for _, value := range nums {
	// 	fmt.Println(value)
	// }

	// fruit := [4]string{"apple", "banana"}

	// fmt.Println(fruit)

	var nums = [...]int{2, 3, 5, 6, 7, 8, 9, 2, 23, 43, 3, 54}

	fmt.Printf("%T\n", nums)
}
