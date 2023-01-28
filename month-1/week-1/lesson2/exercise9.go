package main

import "fmt"

// 9.Berilgan x, y, z sonlari ichidan eng kichigini aniqlash algoritmini tuzing.

func findMinNum(a, b, c int) int {
	var min1 int

	if a < b && a < c {
		min1 = a
	} else if b < a && b < c {
		min1 = b
	} else if c < a && c < b {
		min1 = c
	}

	return min1

}

func main() {
	var x, y, z int = 2, 3, 4

	fmt.Println("Min num from", x, y, z, "is", findMinNum(x, y, z))

}
