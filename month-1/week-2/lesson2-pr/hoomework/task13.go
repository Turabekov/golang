package main

import "fmt"

// Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.

func uniqueOccurrences(arr []int) bool {
	newMap := make(map[int]int)

	for _, v := range arr {
		newMap[v]++
	}

	map2 := make(map[int]int)
	for key, val := range newMap {
		map2[val] = key
	}

	if len(newMap) == len(map2) {
		return true
	} else  {
		return false
	}


}

func main() {
	fmt.Println(uniqueOccurrences([]int{1,2,2,1,1,3}))
}