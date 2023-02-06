package main

import "fmt"

// 4)

func main() {

	x1 := 8
	y1 := 1

	x2 := 3
	y2 := 4

	var text string

	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if x1-1 == x2 || x1-2 == x2 || x1+1 == x2 || x1+2 == x2 {
				if y1-1 == y2 || y1-2 == y2 || y1+1 == y2 || y1+2 == y2 {
					text = "Yura oladi"
				} else {
					text = "Yura olmaydi"
				}
			} else {
				text = "Yura olmaydi"
			}
		}

	}

	fmt.Println(text)

}
