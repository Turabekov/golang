package main

import "fmt"

// Given an array of N integers, and a number sum, the task is to find the number of pairs of integers in the array whose sum is equal to sum.
// Input:  arr[] = {1, 5, 7, -1}, sum = 6
// Output:  2
// Explanation: Pairs with sum 6 are (1, 5) and (7, -1).
// Input:  arr[] = {1, 5, 7, -1, 5}, sum = 6
// Output:  3
// Explanation: Pairs with sum 6 are (1, 5), (7, -1) & (1, 5).

// func sample(arr []int, sum int) int {
// 	var counter = 0

// 	for _, v := range arr {
// 		for i := 0; i < len(arr); i++ {
// 			if v == arr[i] {
// 				continue
// 			}

// 			if v+arr[i] == sum {
// 				counter++
// 			}
// 		}
// 	}

// 	return counter / 2
// }
// O(n*n) 
// ==============================================================
func getPairsCount(arr []int, sum int) int {
	newMap := make(map[int]int)
	count := 0

	for i := 0; i < len(arr); i++ {
		_, ok := newMap[sum - arr[i]]
		if ok {
			count += newMap[sum - arr[i]]
		}

		_, ok2 := newMap[arr[i]]
		if ok2 {
			newMap[arr[i]] = newMap[arr[i]] + 1
		} else {
			newMap[arr[i]] = 1
		}
	}

	return count


}
// O(n)


func main() {
	arr := []int{1, 5, 7, -1}

	fmt.Println(getPairsCount(arr, 6))
}
