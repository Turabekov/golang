package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

// ==================Insert==============================================================
func (l *LinkedList) Insert(data int) {

	newNode := Node{}
	newNode.data = data

	if l.len == 0 {
		l.head = &newNode
		l.len++
		return
	}

	ptr := l.head

	for i := 0; i < l.len; i++ {

		if ptr.next == nil {
			ptr.next = &newNode
			l.len++
			return
		}

		ptr = ptr.next
	}
}

// ==========================InsertAt======================================================

func (l *LinkedList) InsertAt(pos, data int) {

	newNode := Node{}
	newNode.data = data

	if pos < 0 {
		return
	}

	if pos == 0 {
		newNode.next = l.head
		l.head = &newNode
		l.len++
		return
	}

	if pos > l.len {
		return
	}

	n := l.GetAt(pos)

	newNode.next = n

	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}

// ===============================DeletedAt=================================================

func (l *LinkedList) DeletedAt(pos int) {

	if pos < 0 {
		fmt.Println("position can not be negative")
		return
	}

	if l.len == 0 {
		fmt.Println("No nodes in list")
		return
	}

	if pos == 0 {
		l.head = l.head.next
		l.len--
		return
	}

	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		return
	}

	prevNode.next = l.GetAt(pos).next
	l.len--
}

// ============================GetAt====================================================

func (l *LinkedList) GetAt(pos int) *Node {

	if pos > l.len {
		fmt.Println("Not found node")
		return nil
	}

	if pos < 0 {
		return l.head
	}

	ptr := l.head

	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}

	return ptr
}

// =============================Print===================================================

func (l *LinkedList) Print() {

	fmt.Println("LinkedList:")

	ptr := l.head

	for i := 0; i < l.len; i++ {

		fmt.Printf("%d ", ptr.data)
		ptr = ptr.next
	}

	fmt.Println()
}

// =============================Search===================================================
func (l *LinkedList) Search(data int) bool {
	ptr := l.head

	for i := 0; i < l.len; i++ {

		if ptr.data == data {
			return true
		}
		ptr = ptr.next
	}

	return false
}

func main() {

	llist := LinkedList{}

	llist.Insert(40)
	llist.Insert(41)
	llist.Insert(42)
	llist.Insert(43)
	llist.Insert(44)

	llist.InsertAt(2, 54)
	llist.Print()

	llist.DeletedAt(2)
	llist.Print()

	fmt.Println(llist.Search(40))
}
