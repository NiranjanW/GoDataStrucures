package main

import "fmt"

type node struct {
	next *node
	data int
}

type linkedlist struct {
	head   *node
	length int
}

func (l linkedlist) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf(" \n")
}
func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second

	l.length++
}

func (l *linkedlist) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previoustoDelete := l.head
	for previoustoDelete.next.data != value {
		if previoustoDelete.next.next == nil {
			return
		}
		previoustoDelete = previoustoDelete.next
	}
	previoustoDelete.next = previoustoDelete.next.next
	l.length--
}

func main() {
	mylist := linkedlist{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 42}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.printListData()
	mylist.deleteWithValue(42)
	mylist.printListData()
}
