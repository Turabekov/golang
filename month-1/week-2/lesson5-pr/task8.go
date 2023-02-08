package main

import "fmt"

// Given an integer array nums of length n, you want to create an array ans of length 2n where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).
func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := []int{}
	j := 0
	for i := 0; i < 2*n; i++ {

		if i >= n {
			ans = append(ans, nums[j])
			j++
		} else {
			ans = append(ans, nums[i])
		}

	}

	return ans
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1}))
}
