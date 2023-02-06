package main

import "fmt"

func countWord(word string) {
	storage := map[string]int{}

	for _, v := range word {
		storage[string(v)]++
	}

	fmt.Println(storage)

}

func main() {
	var word = "Qo'yho'kiz"
	countWord(word)
}
