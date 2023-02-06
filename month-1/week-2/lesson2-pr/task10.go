package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	arr := []int{}
	newMap := map[int]int{}
	sort.Ints(nums)

	for _, v := range nums {
		// if v  {

		// }
		newMap[v]++
	}

	// for _, val := range nums {
	// 	counter := 0
	// 	for i := 0; i < len(nums); i++ {
	// 		if val == nums[i] {
	// 			continue
	// 		}
	// 		if val > nums[i] {
	// 			counter++
	// 		}
	// 	}
	// 	arr = append(arr, counter)
	// }

	return arr
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))
}
