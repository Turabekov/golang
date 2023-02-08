package main

import "fmt"

func square(nums []int) {
	for ind, val := range nums {
		nums[ind] = val * val
	}
}

func squareMap(nums map[int]int) {
	for key, val := range nums {
		nums[key] = val * val
	}
}

func main() {

	nums := []int{2, 3, 6, 23, 10}
	numsMap := map[int]int {
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
	}

	square(nums)
	squareMap(numsMap)

	fmt.Println(nums)
	fmt.Println(numsMap)
}
