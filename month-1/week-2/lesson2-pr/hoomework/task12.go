package main

import "fmt"


// Given a string s consisting of lowercase English letters, return the first letter to appear twice.

// Note:

// A letter a appears twice before another letter b if the second occurrence of a is before the second occurrence of b.
// s will contain at least one letter that appears twice.


func repeatedCharacter(s string) byte {
	newMap := make(map[string]int)

	for _, val := range s {
		_, ok := newMap[string(val)]
		if ok {
			return byte(val)
		}
		newMap[string(val)]++
	}
	

	return ' '
}	


func main() {
	fmt.Println(string(repeatedCharacter("abccbaacz")))
}