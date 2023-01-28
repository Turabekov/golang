package main

import "fmt"

// 5. Uch xonali son berilgan. Shu sonni palindrom yoki palindrom emasligiga tekshiring.
func main() {
	var number int = 232

	first := number / 100
	// second := number % 100 / 10
	third := number % 100 % 10

	if first == third {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Not Palindrom")

	}

}
