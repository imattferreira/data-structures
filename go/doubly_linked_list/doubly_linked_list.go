package doublylinkedlist

import "fmt"

type Node struct {
	value          int64
	next, previous *Node
}

func NewNode(value int64) *Node {
	return &Node{
		value:    value,
		next:     nil,
		previous: nil,
	}
}

type DoublyLinkedList struct {
	size uint64
	head *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		size: 0,
		head: nil,
	}
}

func (d *DoublyLinkedList) insert(value int64) {
	node := NewNode(value)
	d.size++

	if d.head == nil {
		d.head = node
		return
	}

	tmp := d.head

	for tmp != nil {
		if tmp.next == nil {
			node.previous = tmp
			tmp.next = node
			break
		}

		tmp = tmp.next
	}
}

func (d *DoublyLinkedList) insertAtBeginning(value int64) {
	node := NewNode(value)
	d.size++

	if d.head == nil {
		d.head = node
		return
	}

	node.next = d.head
	d.head.previous = node
	d.head = node
}

func (d *DoublyLinkedList) pop() {
	if d.head == nil {
		return
	}

	d.size--
	tmp := d.head

	for tmp != nil {
		if tmp.next == nil {
			tmp.previous.next = nil
			break
		}

		tmp = tmp.next
	}
}

func (d *DoublyLinkedList) show() {
	tmp := d.head

	for tmp != nil {
		fmt.Println(tmp.value)
		tmp = tmp.next
	}
}
