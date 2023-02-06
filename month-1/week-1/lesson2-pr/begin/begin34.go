package main

import "fmt"

func main() {
	var x, y int = 1, 1
	var A int = 3000
	var B int = 2000

	diff := x*A - y*B
	otnosheniya := float64((x * A)) / float64((y * B))

	fmt.Println(diff, otnosheniya)
}
