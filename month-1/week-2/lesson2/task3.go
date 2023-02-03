package main

import "fmt"

// Given an array, arr[0..n-1] of distinct elements and a range [low, high], find all numbers that are in a range, but not the array. The missing elements should be printed in sorted order.
// Input: arr[] = {10, 12, 11, 15},
//        low = 10, high = 15
// Output: 13, 14

// Input: arr[] = {1, 14, 11, 51, 15},
//        low = 50, high = 55
// Output: 50, 52, 53, 54 55

func returnNotRange(arr []int, low int, high int) []int {
	// newArr := []int{}
	returnArr := []int{}
	newMap := map[int]int{}

	for _, val := range arr {
		newMap[val] = 0
	}

	for i := low; i <= high; i++ {
		_, ok := newMap[i]
		if !ok {
			returnArr = append(returnArr, i)
		}
	}

	return returnArr
}

func main() {
	arr := []int{1, 14, 11, 51, 15}

	newArr := returnNotRange(arr, 50, 55)

	for _, v := range newArr {
		fmt.Printf("%d ", v)
	}

	fmt.Println()

}
