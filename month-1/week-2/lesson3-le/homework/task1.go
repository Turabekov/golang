package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1. Sig'imi 10ta bo'lgan slice e'lon qiling va unga random sonlar soling.
// Shu sonlar ichidan faqat tublarini chiqaruvchi funksiya tuzing.

func Perm(n int) []int {
	m := make([]int, n)

	for i := 0; i < n; i++ {
		j := rand.Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

func getPrimeNum(arr []int) []int {
	storage := []int{}
	count := 0

	for i := 0; i < len(arr); i++ {
		count = 0
		for j := 1; j <= arr[i]; j++ {
			if arr[i]%j == 0 {
				count++
			}
		}
		if count == 2 {
			storage = append(storage, arr[i])
		}
	}

	return storage
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := Perm(10)

	fmt.Println("Random array:", arr)
	fmt.Println("Prime numbers:", getPrimeNum(arr))
}
