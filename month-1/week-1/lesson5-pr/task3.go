package main

import (
	"fmt"
	"math"
)

// 3.Butun qiymat qaytaruvchi DigitN(K, N) funksiyasini hosil qiling. (K > 0).
// Funksiya K sonining N – raqamini qaytarsin. Agar K soni raqamlari N dan kichik bo’lsa,
// minus bir qaytarilsin.

func DigitN(K uint, N int) int {
	num2 := K

	// ==========================
	num3 := 0
	result := 0
	counter := 0
	for K > 0 {
		K = K / 10
		counter++
	}
	// ==========================
	if counter < N {
		return -1
	}

	for num2 > 0 {
		if counter == 0 {
			continue
		}
		num3 += int(num2) % 10 * int(math.Pow(10, float64(counter-1)))
		// fmt.Println("num3 ", num3)
		num2 = num2 / 10
		counter--
	}

	for i := 1; i <= N; i++ {
		result = num3 % 10
		num3 = num3 / 10
	}

	return result
}

func main() {
	fmt.Println(DigitN(12456, 3))
}
