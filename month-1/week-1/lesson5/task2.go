package main

import "fmt"

// 2. x,y haqiqiy sonlari berilgan. Ularning kichigini sonlar yig’indisining yarmiga, kattasini
// ko’paytmasining ikkilanganiga almashtiring. Agar sonlar teng
// bo'lsa, o'zgarishsiz qoldirilsin.

func main() {
	var x, y int = 4, 4

	fmt.Println("Before changing: x =", x, "y =", y)

	if x > y {
		x, y = (x+y)/2, (x*y)/2
	} else if y > x {
		y, x = (x+y)/2, (x*y)/2
	}

	fmt.Println("After changing x =", x, "y =", y)
}
