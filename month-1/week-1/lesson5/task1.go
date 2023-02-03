package main

import "fmt"

// 1.Yil berilgan (musbat butun son). Berilgan yilda nechta kun borligini toping. Kabisa yilida 366 kun bor, kabisa bo’lmagan yilda 365 kun bor.
// Kabisa yil deb 4 ga karrali yillarga aytiladi. Lekin 100 ga karrali yillar ichida faqat 400
// ga karrali bo’lganlari kabisa yil hisoblanadi. Masalan 300, 1300 va 1900 kabisa yili
// emas. 1200 va 2000 kabisa yili.

func isLeapyear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

func main() {
	var year int

	year = 1200

	if isLeapyear(year) {
		fmt.Println("Kabisa yil 366 kun")
	} else {
		fmt.Println("Kabisa yil emas 365 kun")
	}

}
