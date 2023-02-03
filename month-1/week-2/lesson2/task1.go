package main

import "fmt"

// Berilgan massiv palindrom yoki palindrom emasligini aniqlang.
// Input: [1 2 3 2 1]
// Output: true
// Input: [1 2 3 4 2 1]
// Output: false

func isPalindrom(arr []int) bool {
	var newArr = []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		newArr = append(newArr, arr[i])
	}

	for i, v := range newArr {
		if v != arr[i] {
			return false
		}
	}

	return true
}

func isPalindrom2(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		j := len(arr) - 1 - i
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func main() {
	var arr = []int{1, 2, 3, 2, 1}
	// var newMap = map[int]int
	fmt.Println(isPalindrom(arr))
	fmt.Println(isPalindrom2(arr))

}
