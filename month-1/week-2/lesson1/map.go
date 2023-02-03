package main

import "fmt"

func main() {

	// map nil
	// var nums map[int]string = make(map[int]string)
	var nums = map[string]int64 {
		"TWO": 2,
		"ZERO": 0,
	}
	nums["ONE"] = 1

	fmt.Println(nums["TWO"])
	fmt.Println(nums["ONE"])
	val, ok := nums["ZERO"]

	fmt.Println(val, ok)
}
