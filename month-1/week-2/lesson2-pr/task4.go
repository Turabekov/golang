package main

import "fmt"

// Butun sonli massiv berilgan bo‘lsa, massivdagi k-chi unique elementni chiqaring. Berilgan massiv dublikatlarni o'z ichiga olishi mumkin va output barcha unique elementlar orasida k-elementni chop etishi kerak. Agar k unique elementlar sonidan ko'p bo'lsa, -1 ni chiqaring.
// Input : arr[] = {1, 2, 1, 3, 4, 2},
//         k = 2
// Output : 4

// First non-repeating element is 3
// Second non-repeating element is 4

// Input : arr[] = {1, 2, 50, 10, 20, 2},
//         k = 3
// Output : 10

// Input : {2, 2, 2, 2},
//         k = 2
// Output : -1

func returnNonRepeatingElem(arr []int, k int) int {
	obj := map[int]int{}

	for _, v := range arr {
		obj[v]++
	}

	// counter := 0
	fmt.Println(obj)

	for key, v := range obj {
		// if v == 1 {
		// 	fmt.Println("key", key)
		// }

		fmt.Println("key", key, "val", v)
	}

	return -1
}

func main() {
	arr := []int{1, 2, 50, 10, 20, 2}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Println(returnNonRepeatingElem(arr, 3))
}
