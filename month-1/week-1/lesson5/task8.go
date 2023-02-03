package main

import "fmt"

// 8. Berilgan sonning raqamlar yig’indisini shu raqam bir xonali bo’lmaguncha hisoblang.
// Example:
// 	4567  	=>	4+5+6+7=22	 => 2+2=4
// 	123456 =>    1+2+3+4+5+6=21 => 2+1=3

func sumOfDigits(num int) int {
	var sum = 0

	for num > 0 {
		sum += num % 10
		fmt.Println(sum)
		num = num / 10
		if sum/10 != 0 && num == 0 {
			num = sum
			sum = 0
		}
	}

	return sum
}

func main() {
	fmt.Println("Result", sumOfDigits(4567))
}
