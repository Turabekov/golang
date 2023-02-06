package main

import "fmt"

func reverse(nums []int) []int {
	arr := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		arr = append(arr, nums[i])
	}
	return arr
}

func main() {

	var numbers = []int{1, 2, 3, 4, 5, 6}

	reversed := reverse(numbers)

	fmt.Println(reversed) // 6 5 4 3 2 1

}
