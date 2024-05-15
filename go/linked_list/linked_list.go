package linkedlist

import "fmt"

type Node struct {
	value int64
	next  *Node
}

func NewNode(value int64) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

type LinkedList struct {
	size uint64
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		size: 0,
		head: nil,
	}
}

func (ll *LinkedList) insert(value int64) {
	node := NewNode(value)
	ll.size++

	if ll.head == nil {
		ll.head = node
		return
	}

	tmp := ll.head

	for tmp != nil {
		if tmp.next == nil {
			tmp.next = node
			break
		}

		tmp = tmp.next
	}
}

func (ll *LinkedList) insertAtBeginning(value int64) {
	node := NewNode(value)
	ll.size++

	node.next = ll.head
	ll.head = node
}

func (ll *LinkedList) pop() {
	if ll.head == nil {
		return
	}

	var previous *Node = nil
	current := ll.head
	ll.size--

	for current != nil {
		if current.next == nil {
			previous.next = nil
			break
		}

		previous = current
		current = current.next
	}
}

func (ll *LinkedList) show() {
	node := ll.head

	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}
