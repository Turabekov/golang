package main

import "fmt"

func findMin(sliceNums []int) int {

	minimum := sliceNums[0]

	for i := 1; i < len(sliceNums); i++ {
		if sliceNums[i] < minimum {
			minimum = sliceNums[i]
		}
	}

	return minimum
}

func main() {

	var slice []int = []int{5, 2, 6, 7}

	fmt.Printf("Minimum among slice %d is %d", slice, findMin(slice))
}
