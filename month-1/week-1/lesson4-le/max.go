package main

import "fmt"

// variadic function
func max(nums ...int) int {

	if !(len(nums) > 0) {
		return 0
	}

	maximum := nums[0]

	for _, val := range nums {
		
		if val > maximum {
			maximum = val
		}
	}

	return maximum
}

// sum(3, 4, 5, 6, 7, 2, 10, 213, 543)

func main() {

	var nums = []int{3, 4, 5, 6}
	var sl = []int{7, 8, 9}

	sl = append(sl, nums...)

	fmt.Println(sl)

	// fmt.Println(max(nums...))


}
