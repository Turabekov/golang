package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	fmt.Println(nums)

	// 1. Random
	// nums [500]int

	// juft []int
	// toq []int

	// fmt.Println(juft)
	// fmt.Println(toq)
}
