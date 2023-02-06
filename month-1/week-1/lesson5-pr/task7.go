package main

import "fmt"

// 7. Berilgan slice da minimal elementlar sonini toping.
// Example:
// Input:	[1 4 8 9 1 1 2 3 4]
// Output: 3

func findMinimum(nums []int) int {
	minimum := nums[0]

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		if minimum > nums[i] {
			minimum = nums[i]
		}
	}

	counter := 0
	for _, val := range nums {
		if minimum == val {
			counter++
		}
	}

	return counter
}

func main() {
	nums := []int{1, 4, 8, 9, 1, 1, 2, 3, 4, 1}

	fmt.Println(findMinimum(nums))
}
