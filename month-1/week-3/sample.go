package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	ptr := head

	counter := 0
	for ptr == nil {
		counter++
		ptr = ptr.Next
	}

	fmt.Println(counter)

	return head

}

func main() {
	// fmt.Println(reverseList())
}
