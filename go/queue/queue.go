package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	capacity, size uint64
	items          []*int64
	front, tail    *int64
}

func NewQueue(capacity uint64) *Queue {
	return &Queue{
		capacity: capacity,
		size:     0,
		items:    []*int64{},
		front:    nil,
		tail:     nil,
	}
}

func (q *Queue) clear() {
	// TODO: finish implementation (wrong implementation is here) -> FIFO
	// for q.size != 0 {
	// 	q.items[q.size-1] = nil
	// 	q.size--
	// }

	// q.front = nil
	// q.tail = nil
}

func (q *Queue) dequeue() {
	// TODO: finish implementation (wrong implementation is here) -> FIFO
	// if q.empty() {
	// 	return
	// }

	// q.size--
	// q.items[q.size] = nil
	// q.tail = q.items[q.size-1]
}

func (q *Queue) enqueue(item *int64) error {
	if q.full() {
		return errors.New("[queue] is full!")
	}

	q.items[q.size] = item
	q.tail = item
	q.size++

	return nil
}

func (q *Queue) empty() bool {
	return q.size == 0
}

func (q *Queue) full() bool {
	return q.size == q.capacity
}

func (q *Queue) show() {
	var i uint64 = 0

	for i < q.size {
		fmt.Println(q.items[i])
		i++
	}
}
