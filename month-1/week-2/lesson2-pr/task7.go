package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	newMap := map[int]int{}
	counter := 0

	for _, num := range nums {
		counter = counter + newMap[num]
		newMap[num]++
	}

	return counter

}

func main() {

	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))

}
