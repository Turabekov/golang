package main

import (
	"fmt"
)

func main() {

	// var company string

	// company = "Apple"

	// fmt.Printf("%T\n", company)
	// fmt.Println(company)

	// fmt.Println(company[:2])
	// fmt.Println(company[2:])
	// fmt.Println(company[2:3])
	// fmt.Println(string(company[len(company) - 1]))

	// for _, val := range company {
	// 	fmt.Printf("%c\n", val)
	// }

	var word = "Qo'yho'kiz"

	// countWord(word)

	// Q 1
	// o 2
	// ' 2
	// h 1
	// k 1
	// i 1
	// z 1

	nums := map[string]int{}

	for _, element := range word {
		nums[string(element)]++
	}

	for key, val := range nums {
		fmt.Println(key, val)
	}

	// 2. 6543212 = 2 - raqamini: 5

	// var num = 6543212

	// fmt.Println(string(fmt.Sprintf("%d", num)[1]))

	// numStr := strconv.Itoa(num)

	// if len(numStr) > 1 {
	// 	fmt.Println(numStr[1])
	// }
}
