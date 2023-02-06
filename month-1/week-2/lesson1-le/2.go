package main

import (
	"fmt"
)

func main() {

	// var status map[string]bool = map[string]bool {
	// 	"access": true,
	// 	"failed": true,
	// 	"new": true,
	// }

	// delete(status, "new")

	// // fmt.Println(status)

	// for key, val := range status {
	// 	fmt.Println(key, val)
	// }

	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 2, 3, 4, 2}
	nums := map[int]int{}

	for _, element := range sl {
		nums[element] = nums[element] + 1
	}

	for key, val := range nums {
		fmt.Println(key, val)
	}
}