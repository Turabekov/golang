package main

import (
	"fmt"
	"strconv"
)

// func printNum(num int, index int) int {
// 	num2 := num
// 	num3 := 0
// 	result := 0
// 	counter := 0
// 	for num > 0 {
// 		num = num / 10
// 		counter++
// 	}

// 	for num2 > 0 {
// 		if counter == 0 {
// 			continue
// 		}
// 		num3 += num2 % 10 * int(math.Pow(10, float64(counter-1)))
// 		// fmt.Println("num3 ", num3)
// 		num2 = num2 / 10
// 		counter--
// 	}

// 	for i := 1; i <= index; i++ {
// 		result = num3 % 10
// 		num3 = num3 / 10
// 	}

// 	fmt.Println(result)

// 	return result

// }

func printNum(num int, index int) {
	str := strconv.Itoa(num)

	for i, v := range str {
		if i+1 == index {
			fmt.Println(string(v))
		}
	}
}

func main() {
	printNum(6543212, 4)
}
