package main

import (
	"fmt"
	"sort"
)

func bubbleSort(nums []int) {
	var count int

	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}

			count++
		}
	}

	fmt.Println(nums)
	fmt.Println(count)
}

func main() {
	var nums = []int{8, 5, 6, 12, 32, 48, 129, 23, 58}

	// bubbleSort(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums)
}
