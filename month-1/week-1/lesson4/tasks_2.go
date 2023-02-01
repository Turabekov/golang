package main

import "fmt"

// 1)
func changeOrder(num int) int {
	a := num / 10
	b := num % 10

	newNum := b*10 + a
	return newNum
}

// 2)
func lenOfNum(num int) int {
	counter := 0
	for num > 0 {
		fmt.Println(num)
		num = num / 10
		counter++
	}

	return counter
}

// 3)
func sumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		fmt.Println("Ostatok: ", num%10)
		fmt.Println("num", num)

		sum += num % 10
		num = num / 10
	}
	return sum
}

// 5)
func checkPalindrom(num int) string {
	input_num := num
	var remainder int
	res := 0

	for num > 0 {
		remainder = num % 10
		res = (res * 10) + remainder
		num = num / 10
	}

	if input_num == res {
		return "Palindrome"
	} else {
		return "Not a Palindrome"
	}
}

func main() {

	// fmt.Println(factorial(5))
	// fmt.Println(sum(23, 55, 22))
	// fmt.Println(changeOrder(24))
	// fmt.Println(lenOfNum(444554))
	// fmt.Println(sumOfDigits(123))
	// fmt.Println(checkPalindrom(325))

}
