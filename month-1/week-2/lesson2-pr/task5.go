package main

import "fmt"

// Unique elementlarga ega massiv berilgan boʻlsa, a % b = k boʻladigan juftliklarni sonini toping.
// Input  :  arr[] = {2, 3, 5, 4, 7}
//               k = 3
// Output :  (7, 4), (3, 4), (3, 5), (3, 7)
// 7 % 4 = 3
// 3 % 4 = 3
// 3 % 5 = 3
// 3 % 7 = 3

func sample(arr []int, k int) {
	for _, v := range arr {
		for i := 0; i < len(arr); i++ {
			if arr[i] == v {
				continue
			}
			if v%arr[i] == 3 {
				fmt.Printf("(%d %d) \t", v, arr[i])
			}
		}
	}

}
// O(n*n)
// ==============================================================
// func sample(arr []int, k int) {
// 	newMap := make(map[int][]int)
// 	for i := 0; i < len(arr); i++ {
		
// 	}

// }
O(n)
//====================================================
func main() {
	arr := []int{2, 3, 5, 4, 7}

	sample(arr, 3)
}
