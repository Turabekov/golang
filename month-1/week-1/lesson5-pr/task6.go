package main

import "fmt"

// 6.N natural soni berilgan. Shu songacha bo’lgan tub sonlarni chiqaring.

// 10 			2,3,5,7
// 13			2,3,5,7,11,13
type Ostatok []int

func findPrimeNums(num int) []int {
	dbNumbers := []int{}
	mapOfNums := map[int]Ostatok{}

	for i := 1; i <= num; i++ {
		dbDelitel := []int{}

		for j := 1; j <= i; j++ {
			if i%j == 0 {
				dbDelitel = append(dbDelitel, j)
			}
		}
		mapOfNums[i] = dbDelitel
	}

	for key, ostatokSlice := range mapOfNums {
		if len(ostatokSlice) == 2 {
			dbNumbers = append(dbNumbers, key)
		}
	}

	return dbNumbers
}

func main() {
	nums := findPrimeNums(13)
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
