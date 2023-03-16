package main

import (
	"fmt"
	"strconv"
)

// Raqamlar qatorini hisobga olgan holda, 5 dan past bo'lgan har qanday raqamni "0" bilan, 5 va undan yuqori har qanday raqamni "1" bilan almashtirishingiz kerak. Olingan qatorni qaytaring
func changeStr(str string) string { 
	ans := ""
	for _, c := range str {
		d, err := strconv.Atoi(string(c))
		if err == nil {
			if d < 5 {
				ans += "0"
			} else if d >= 5 {
				ans += "1"
			}
		} else {
			ans += string(c)
		}
	}

	return ans
}

func main() {
	fmt.Println(changeStr("23asddsf6785"))
}
