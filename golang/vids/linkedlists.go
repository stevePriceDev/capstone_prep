/*
-linked list - like an array but with nodes instead of values
-nodes are linked together - each node contains an address to next node
-to find node, go to head and follow along nodes until found
-adding node to beginning is faster than array
	O(1)
-adding value to beginning of array must move every value over
	O(N)
-doubly linked list - nodes contain address to next and previous nodes
*/
package main
import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}

func (l *linkedList) prepend (n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList{}
	node1 := &node{data:48}
	node2 := &node{data:18}
	node3 := &node{data:16}
	node4 := &node{data:11}
	node5 := &node{data:7}
	node6 := &node{data:2}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.printListData()
	mylist.deleteWithValue(16)
	mylist.printListData()
	mylist.deleteWithValue(2)
	mylist.printListData()
}

