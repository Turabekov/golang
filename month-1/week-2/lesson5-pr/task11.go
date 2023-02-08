package main

import "fmt"

// A sentence is a list of words that are separated by a single space with no leading or trailing spaces.

// You are given an array of strings sentences, where each sentences[i] represents a single sentence.

// Return the maximum number of words that appear in a single sentence.

func mostWordsFound(sentences []string) int {
	newMap := map[int]int{}

	for i := 0; i < len(sentences); i++ {
		countWord := 0
		for j := 0; j < len(sentences[i]); j++ {
			val := sentences[i]
			if string(val[j]) == " " {
				countWord++
			}
		}
		newMap[i] = countWord + 1
	}

	max := newMap[0]
	for _, val := range newMap {
		if max < val {
			max = val
		}
	}

	return max
}

func main() {
	fmt.Println(mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))
}
