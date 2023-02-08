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

// 1. Linked list uchun addFirst() metodini yozing.
// =====================AddFirst================================================================================
func (l *LinkedList) addFirst(data int) {
	newNode := Node{}
	newNode.data = data

	if l.len == 0 {
		l.head = &newNode
		l.len++
		return
	}

	ptr := l.head
	l.head = &newNode
	newNode.next = ptr
	l.len++
}

// 2. Linked list uchun removeFirst() metodini yozing
// =====================removeFirst================================================================================
func (l *LinkedList) removeFirst() {

	if l.len == 0 {
		fmt.Println("No nodes in list")
		return
	}

	l.head = l.head.next
	l.len--
}

// 3. Linked listni berilgan n-position dagi elementini qaytaring
// =====================getPosElem()================================================================================
func (l *LinkedList) getPosElem(n int) int {

	if l.len == 0 {
		fmt.Println("No nodes in list")
		return 0
	}

	if n > l.len {
		fmt.Println("n out if range")
		return 0
	}

	ptr := l.head

	for i := 0; i < n; i++ {
		ptr = ptr.next
	}
	fmt.Printf("Elem in %d pos is: %d\n", n, ptr.data)

	return ptr.data
}

// 4. Linked listni middle elementini toping.
// =====================getMiddleElem()================================================================================
func (l *LinkedList) getMiddleElem() int {

	if l.len == 0 {
		fmt.Println("No nodes in list")
		return 0
	}

	ptr := l.head

	for i := 0; i < l.len/2; i++ {
		ptr = ptr.next
	}
	fmt.Printf("Middle elem is: %d\n", ptr.data)

	return ptr.data

}

// 5. Linked listni middle elementini o’chiring.
// =====================removeMiddleElem()================================================================================
func (l *LinkedList) removeMiddleElem() {

	if l.len == 0 {
		fmt.Println("No nodes in list")
		return
	}

	ptr := l.head

	for i := 0; i < l.len/2; i++ {
		ptr = ptr.next
	}
	prev := l.head
	for i := 0; i < l.len/2-1; i++ {
		prev = prev.next
	}

	prev.next = ptr.next
	l.len--
}

// 6. Linked listni oxirgi elementidan boshlab n-positiondagi elementini qaytaring.
// =====================getReverseNElem()================================================================================
func (l *LinkedList) getReverseNElem(n int) int {

	if l.len == 0 {
		fmt.Println("No nodes in list")
		return 0
	}

	ptr := l.head

	for i := 0; i < l.len-n; i++ {
		ptr = ptr.next
	}

	fmt.Println("N elem from last is", ptr.data)
	return ptr.data

}

// 7. Linked listni birinchi va oxirgi elementlarini valuelarini almashtiring.
// =====================changeFirstAndLast()================================================================================
func (l *LinkedList) changeFirstAndLast() {

	if l.len == 0 {
		fmt.Println("No nodes in list")
		return
	}

	firstElem := l.head
	ptr := l.head

	for i := 0; i < l.len-1; i++ {
		ptr = ptr.next
	}

	firstElem.data, ptr.data = ptr.data, firstElem.data
}

// ===================Print============================
func (l *LinkedList) Print() {

	fmt.Println("LinkedList:")

	ptr := l.head

	for i := 0; i < l.len; i++ {

		fmt.Printf("%d ", ptr.data)
		ptr = ptr.next
	}

	fmt.Println()
}

// ===============================================
func main() {
	linklist := LinkedList{}

	linklist.addFirst(10)
	linklist.addFirst(20)
	linklist.addFirst(30)
	linklist.addFirst(40)
	linklist.addFirst(50)
	linklist.addFirst(60)
	linklist.addFirst(70)
	// linklist.addFirst(80)

	linklist.Print()
	// linklist.removeFirst()
	fmt.Println("Removed first elem")
	linklist.Print()
	linklist.getPosElem(1)
	linklist.getMiddleElem()
	linklist.removeMiddleElem()
	linklist.getReverseNElem(2)
	linklist.Print()
	linklist.changeFirstAndLast()
	linklist.Print()

}
