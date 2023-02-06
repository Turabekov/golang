package main

import "fmt"

func main() {

	// var nums []int

	// fmt.Println("Len: ", len(nums), "Cap:", cap(nums))
	// nums = append(nums, 7)
	// fmt.Println("Len: ", len(nums), "Cap:", cap(nums))
	// nums = append(nums, 8)
	// fmt.Println("Len: ", len(nums), "Cap:", cap(nums))
	// nums = append(nums, 8)
	// fmt.Println("Len: ", len(nums), "Cap:", cap(nums))
	// nums = append(nums, 8)
	// fmt.Println("Len: ", len(nums), "Cap:", cap(nums))
	// nums = append(nums, 8)
	// fmt.Println("Len: ", len(nums), "Cap:", cap(nums))

	// var nums []int = make([]int, 10, 11)

	// fmt.Println("Len:", len(nums), "Cap:", cap(nums))
	// fmt.Printf("%T\n", nums)
	// fmt.Println(nums)

	var nums []int = []int{1, 2, 4, 5, 6}
	k := make([]int, 6)

	n := copy(k, nums)

	k[0] = 10

	fmt.Println(k, n)
	fmt.Println(nums)
}