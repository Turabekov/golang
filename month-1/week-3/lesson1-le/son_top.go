package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	size := 10

	var (
		random_number = rand.Intn(size)
		input         int
	)
	chance := math.Log2(float64(size))

	fmt.Println(random_number)

	fmt.Printf("Sizda %d imkoniyat bor sonni topishga\n", int(chance))
	for {
		fmt.Println("Son kiriting:")

		fmt.Scan(&input)

		if chance <= 0 {

			fmt.Println("Siz yutqazdingiz 😛")
			break
		}

		if input == random_number {
			fmt.Println("Topildi 😀😀😀")
			break
		} else {
			if input < random_number {
				fmt.Println("Kiritgan soningiz kichikroq😀")
			} else if input > random_number {
				fmt.Println("Kiritgan soningiz kattaroq😀")
			}

			chance--
			fmt.Printf("Topilmadi 😔. Imkoniyatingiz %d qoldi.\n", int(chance))
		}
	}
}
