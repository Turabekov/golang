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

func getKElem(arr []int, k int) int {
	newMap := make(map[int]int)

	for i := 0; i <  len(arr); i++ {
		newMap[arr[i]]++
	}

	if len(newMap) < k {
		return -1		
	}

	counter := 0
	for i := 0; i <  len(arr); i++ {
		if newMap[arr[i]] == 1 {
			counter++
		}
		if counter == k {
			return arr[i]
		}
	}

	return -1

}

// ===========================================================================================

func main() {
	arr := []int{1, 2, 50, 10, 20, 2}

	fmt.Println(getKElem(arr, 2))
}
