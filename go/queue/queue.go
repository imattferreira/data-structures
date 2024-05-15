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
	var i uint64 = 0

	for i < (q.size - 1) {
		q.items[i] = nil
		i++
	}

	q.front = nil
	q.tail = nil
	q.size = 0
}

func (q *Queue) dequeue() {
	if q.empty() {
		return
	}

	var i uint64 = 1

	for i <= q.size {
		q.items[i-1] = q.items[i]
		i++
	}

	q.front = q.items[0]
	q.size--
}

func (q *Queue) enqueue(item *int64) error {
	if q.full() {
		return errors.New("[queue] is full")
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
