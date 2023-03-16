package main

import "fmt"

func indexesOfDuplicate(nums []int) []int {
	hash := map[int]int{}
	arr := []int{}

	for i, num := range nums {
		if _, ok := hash[num]; ok {
			arr = append(arr, hash[num])
			arr = append(arr, i)
		}
		hash[num] = i
	}
	return arr
}

func removeDuplicates(nums []int) []int {
	i := 0
	j := 0
	for j != len(nums) {
		if nums[i] != nums[j] {
			nums[i+1], nums[j] = nums[j], nums[i+1]
			i++
		}
		j++
	}
	return nums[0 : i+1]
}

func main() {
	fmt.Println(indexesOfDuplicate([]int{1, 2, 2, 3, 3, 4, 5}))
	fmt.Println(removeDuplicates([]int{1, 2, 2, 3, 3, 4, 5}))
}
