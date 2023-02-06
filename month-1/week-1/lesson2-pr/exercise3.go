package main

import "fmt"

// 3. Uchburchakning 3 ta tomoni berilgan. Berilgan tomonlar orqali  uchburchakning turini aniqlang.

func findMinNum(a, b, c int) (int, int) {
	var min1 int
	var min2 int

	if a < b && a < c {
		min1 = a
		if b < c {
			min2 = b
		} else {
			min2 = c
		}
	} else if b < a && b < c {
		min1 = b
		if a < c {
			min2 = a
		} else {
			min2 = c
		}
	} else if c < a && c < b {
		min1 = c
		if a < b {
			min2 = a
		} else {
			min2 = b
		}
	}

	return min1, min2

}

func findMaxNum(a, b, c int) int {
	var max int
	if a > b && a > c {
		max = a
	} else if b > a && b > c {
		max = b
	} else {
		max = c
	}

	return max
}

func main() {
	var a, b, c int = 5, 3, 4
	side1, side2 := findMinNum(a, b, c)
	maxSide := findMaxNum(a, b, c)

	squareSides := side1*side1 + side2*side2

	if maxSide*maxSide == squareSides {
		fmt.Println("треугольник прямоугольный")
	} else if maxSide*maxSide > squareSides {
		fmt.Println("треугольник тупоугольный.")
	} else {
		fmt.Println("треугольник остроугольный")
	}

}
