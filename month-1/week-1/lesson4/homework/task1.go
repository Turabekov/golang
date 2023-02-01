package main

import "fmt"

func findMin(sliceNums []int) int {
	if len(sliceNums) == 0 {
		return 0
	}

	minimum := sliceNums[0]

	for i := 1; i < len(sliceNums); i++ {
		if sliceNums[i] < minimum {
			minimum = sliceNums[i]
		}
	}

	return minimum
}

func main() {

	var slice []int = []int{}

	fmt.Printf("Minimum number among slice %d is %d", slice, findMin(slice))
}
