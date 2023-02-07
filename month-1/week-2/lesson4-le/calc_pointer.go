package main

import "fmt"

// =================Calc============================
func Add(a, b, res *int) {
	*res = *a + *b
}
func Sub(a, b, res *int) {
	*res = *a - *b
}
func Div(a, b, res *int) {
	*res = *a / *b
}
func Mult(a, b, res *int) {
	*res = *a * *b
}

// =================================================
func main() {
	var x, y int = 10, 5
	var res int

	Add(&x, &y, &res)
	fmt.Println(res)
	Sub(&x, &y, &res)
	fmt.Println(res)
	Div(&x, &y, &res)
	fmt.Println(res)
	Mult(&x, &y, &res)
	fmt.Println(res)

}
