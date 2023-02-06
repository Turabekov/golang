package main

import "fmt"

// 4.Ikkita butun son berilgan Day (kun) va Month (oy). (Kabisa bo`lmagan yil sanasi
// kiritiladi). Berilgan sanadan keyingi sanani chiqaring.

// 7 7 		 08.07
// 31 1		 01.02
// 31 12		 01.01
// 31 15	 	 Bunday oy yo'q
// 35 3 		 Bunday sana yo'q
func main() {
	var day uint8 = 35
	var month uint8 = 3
	var dayError string = ""
	var monthError string = ""

	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		if day < 31 {
			day++
		} else if day == 31 {
			if month == 12 {
				month = 1
			} else {
				month++
			}
			day = 01
		} else {
			dayError = "Bunday sana yo'q"
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		if day < 30 {
			day++
		} else if day == 30 {
			day = 01
			month++
		} else {
			dayError = "Bunday sana yo'q"
		}
	} else if month == 2 {
		if day < 29 {
			day++
		} else if day == 29 {
			day = 01
			month++

		} else {
			dayError = "Bunday sana yo'q"
		}
	} else {
		monthError = "Bunday oy yoq"
	}

	// check exist of error
	if dayError != "" {
		fmt.Println(dayError)
	} else if monthError != "" {
		fmt.Println(monthError)
	} else {
		if day <= 10 {
			fmt.Printf("0%d.0%d\n", day, month)
		} else {
			fmt.Printf("%d.%d\n", day, month)
		}
	}

}
