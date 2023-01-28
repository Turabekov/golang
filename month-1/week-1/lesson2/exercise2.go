package main

import "fmt"

func main() {
	var n int = 5

	for i := 1; true; i++ {
		if i%2 == 0 && i%n == 0 {
			fmt.Println(i)
			break
		}
	}

}
