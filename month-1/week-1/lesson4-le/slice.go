package main

import "fmt"

func main() {

	var nums []int

	nums = append(nums, 4, 5, 6, 7, 8)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)

	fmt.Println(nums)
	fmt.Println(nums[1:2])

	for index, val := range nums {
		fmt.Println(index, val)
	}

	
}
