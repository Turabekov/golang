package main

import "fmt"

func main() {

	var dictionary = map[string]string{
		"apple":  "olma",
		"banana": "banan",
		"pear":   "nok",
		"qiwi":   "kiwi",
		"orange": "apelsin",
	}

	fmt.Println(dictionary)

	fmt.Println(dictionary["apple"])
	fmt.Println(dictionary["banana"])
	fmt.Println(dictionary["pear"])
	fmt.Println(dictionary["qiwi"])
	fmt.Println(dictionary["orange"])

	var word string
	if _, ok := dictionary[word]; !ok {
		fmt.Println("Not found")
		return
	}

}
