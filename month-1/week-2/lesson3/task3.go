package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	rangeNum := max - min
	return rand.Intn(rangeNum) + min
}

func main() {
	nums := []int{}
	juft := []int{}
	toq := []int{}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 500; i++ {
		n := Random(-250, 250)
		nums = append(nums, n)
	}

	for _, v := range nums {
		if v%2 == 0 {
			juft = append(juft, v)
		} else {
			toq = append(toq, v)
		}
	}

	fmt.Println("Juft", juft)
	fmt.Println("Toq", toq)
}
