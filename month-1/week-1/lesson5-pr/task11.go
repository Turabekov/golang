package main

import "fmt"

// 11.N natural soni va slice berilgan.Slice ichidagi eng kichik musbat sonni aniqlang. Agar musbat son bo’lmasa nol chiqarilsin.
// Input: [-1 -2 -9 4 3 1 2 10]
// Output:1

// Input:[-1 -2 -9 -4 -3]
// Output:0

func findMinimumNum(nums []int) int {
	min := 0

	for _, val := range nums {
		if val < 0 {
			continue
		}
		min = val
		break
	}

	for _, v := range nums {
		if v < 0 {
			continue
		}
		if min >= v {
			min = v
		}
	}

	if min == 0 {
		return 0
	}

	return min
}

func main() {
	// time O(n)
	fmt.Println(findMinimumNum([]int{-1, -2, -9, -4, -3, 9}))
}
