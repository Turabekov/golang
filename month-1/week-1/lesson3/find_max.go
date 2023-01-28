package main

import "fmt"

var output = "Result: max num is"

func findMaxNum(a, b, c int) (int, int, int, int) {
	var max int
	if a > b && a > c {
		max = a
	} else if b > a && b > c {
		max = b
	} else {
		max = c
	}

	return max, a, b, c
}

func main() {
	maxN, a, b, c := findMaxNum(7, 5, 6)

	fmt.Printf("%s %d from %d %d %d numbers\n", output, maxN, a, b, c)
}
