package main

import "fmt"

func isAnagram(s string, t string) bool {
	map1 := make(map[string]int)
	map2 := make(map[string]int)

	if len(s) != len(t) {
		return false
	}

	for _, v := range s {
		map1[string(v)]++
	}
	for _, v := range t {
		map2[string(v)]++
	}
	fmt.Println(map1)
	fmt.Println(map2)

	for key, val1 := range map1 {
		val2, ok := map2[string(key)]
		if !ok {
			return false
		}

		if val1 != val2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("aacc", "ccac"))
}
