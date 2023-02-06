package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	newMap := map[string]int{}
	counter := 0

	for _, value := range jewels {
		newMap[string(value)]++
	}

	for _, v := range stones {
		_, ok := newMap[string(v)]
		if ok {
			counter++
		}
	}

	fmt.Println(newMap)

	return counter

}

func main() {
	fmt.Println(numJewelsInStones("z", "ZZ"))
}
