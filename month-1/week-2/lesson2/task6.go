package main

import "fmt"

// Given an array of N integers, and a number sum, the task is to find the number of pairs of integers in the array whose sum is equal to sum.
// Input:  arr[] = {1, 5, 7, -1}, sum = 6
// Output:  2
// Explanation: Pairs with sum 6 are (1, 5) and (7, -1).
// Input:  arr[] = {1, 5, 7, -1, 5}, sum = 6
// Output:  3
// Explanation: Pairs with sum 6 are (1, 5), (7, -1) & (1, 5).

func sample(arr []int, sum int) int {
	var counter = 0

	for _, v := range arr {
		for i := 0; i < len(arr); i++ {
			if v == arr[i] {
				continue
			}

			if v+arr[i] == sum {
				counter++
			}
		}
	}

	return counter / 2
}

func main() {
	arr := []int{1, 5, 7, -1, 5}

	fmt.Println(sample(arr, 6))
}
