package main

import "fmt"

// 5.N natural soni berilgan. Shu songacha bo’lgan mukammal sonlarni chiqaring. O`zidan boshqa bo’luvchilari yig’indisi o’ziga teng bo’lgan son
// mukammal son deyiladi. Masalan: 6, 28

// 100 			6 = 3 + 2 + 1
// 28

// 10000 			6
// 28
// 496
// 8128

type Ostatok []int

func findOriginalNum(num int) []int {
	dbNumbers := []int{}
	mapOfNums := map[int]Ostatok{}

	for i := 1; i <= num; i++ {
		dbDelitel := []int{}

		for j := 1; j < i; j++ {
			if i%j == 0 {
				dbDelitel = append(dbDelitel, j)
			}
		}
		mapOfNums[i] = dbDelitel
	}
	sum := 0

	for key, ostatokSlice := range mapOfNums {
		sum = 0
		for i := 0; i < len(mapOfNums[key]); i++ {
			sum += ostatokSlice[i]
		}
		if key == sum {
			dbNumbers = append(dbNumbers, key)
		}
	}

	return dbNumbers
}

func main() {
	nums := findOriginalNum(10000)
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
