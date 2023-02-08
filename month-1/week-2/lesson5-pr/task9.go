package main

import "fmt"

// Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].

// Return the array in the form [x1,y1,x2,y2,...,xn,yn].
func shuffle(nums []int, n int) []int {
	arrX := []int{}
	arrY := []int{}
	arr := []int{}

	for i := 0; i < n; i++ {
		arrX = append(arrX, nums[i])
	}

	for i := n; i < len(nums); i++ {
		arrY = append(arrY, nums[i])
	}

	for i := 0; i < n; i++ {
		arr = append(arr, arrX[i])
		arr = append(arr, arrY[i])

	}

	return arr
}
func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
