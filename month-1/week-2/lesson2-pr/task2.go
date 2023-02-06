package main

import "fmt"

// Slicedagi eng ko’p marta takrorlangan elementni qaytaring. Agar bunday elementlar bir necta bo’lsa unda eng kichigini qaytaring.
// Input:[]int{1,1,4,5,3,3,3,3,1,2}
// Output:3

func findRepeatElem(nums []int) int {
	newMap := map[int]int{}
	result := 0

	for _, v := range nums {
		newMap[v]++
	}

	max := 0
	for key, v := range newMap {
		if max < v {
			max = v
			result = key
		} else if max == v {
			if result > key {
				result = key
			}
		}
	}

	fmt.Println(newMap)

	return result
}

func main() {
	var nums = []int{1, 1, 4, 5, 3, 3, 3, 3, 1, 2, 2, 2, 2, 1}
	fmt.Println(findRepeatElem(nums))
}
