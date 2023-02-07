package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	
	newMap := make(map[int]int)

	for i := 0; i < len(sortedNums); i++ {
		_, ok := newMap[sortedNums[i]]
		if !ok {
			newMap[sortedNums[i]] = i
		}

	}

	arr := make([]int, len(nums))
	
	for i, val := range nums {
		arr[i] = newMap[val]
	}

	return arr
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8,1,2,2,3}))
}
