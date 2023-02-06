package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 2, 3, 4, 2}

	newMap := map[int]int{}

	for _, v := range nums {
		value, ok := newMap[v]
		if ok {
			value++
			newMap[v] = value
		} else {
			newMap[v] = 1
		}
	}

	fmt.Println(newMap)

}
