package main

import "fmt"

// 7. A ,B,C sonlari berilgan. A ni qiymati C ga , C ni qiymati B ga va B ni qiymatini A ga almashtiring.

func main() {
	var A, B, C int = 1, 2, 3

	fmt.Printf("Before changing: A = %d B = %d C = %d\n", A, B, C)

	A, B, C = C, A, B

	fmt.Printf("After changing: A = %d B = %d C = %d\n", A, B, C)
}
