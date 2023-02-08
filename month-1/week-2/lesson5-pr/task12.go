package main

import (
	"fmt"
)

func restoreString(s string, indices []int) string {
	newMap := make(map[int]string)

	for i := 0; i < len(indices); i++ {
		newMap[indices[i]] = string(s[i])
	}

	var word string
	for i := 0; i < len(indices); i++ {
		word += newMap[i]
	}

	fmt.Println(word)

	return word
}

func main() {
	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
}
