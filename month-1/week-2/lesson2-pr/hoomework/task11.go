package main

import "fmt"

func sumOfUnique(nums []int) int {
	newMap := map[int]int{}

	for _, val := range nums {
		newMap[val]++
	}

	sum := 0
	for key, val := range newMap {
		if val == 1 {
			sum += key
		}
	}

	if sum != 0 {
		return sum
	}

	return 0
}

func main() {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))
}
