package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack_LinkedList struct {
	head *Node
	len int
}

// pushing element in first
func (s *Stack_LinkedList) push(data int) {
	newNode := Node{}
	newNode.data = data

	if s.len == 0 {
		s.head = &newNode
		s.len++
		return
	}

	ptr := s.head
	s.head = &newNode
	s.head.next = ptr
	s.len++
}
// removing element in first position
func (s *Stack_LinkedList) pop() {

	if s.len == 0 {
		fmt.Println("No data for deleting")
		return
	}

	ptr := s.head
	s.head = ptr.next
	s.len--
}

// clear stack
func (s *Stack_LinkedList) clear() {
	if s.len == 0 {
		fmt.Println("No data for clearing")
		return
	}

	n := s.len
	for i := 0; i < n; i++ {
		s.pop()
	}

}


// print stack
func (s *Stack_LinkedList) print() {
	if s.len == 0 {
		fmt.Println("No data for printing")
		return
	}
	fmt.Println("Stack linkedList:")
	
	ptr := s.head

	for i := 0; i < s.len; i++ {
		fmt.Printf("%d ", ptr.data)
		ptr = ptr.next
	}

	fmt.Println()
}


func main(){
	stack := Stack_LinkedList{}

	// push data
	stack.push(5)
	stack.push(6)
	stack.push(7)
	stack.push(8)
	stack.push(9)
	stack.push(10)
	stack.print()
	// pop
	fmt.Println("After deleted first elem")
	stack.pop()
	stack.print()
	// clear
	stack.clear()
	fmt.Println("After cleared all stack")
	stack.print()


	stack.push(5)
	stack.push(6)
	stack.print()


}	