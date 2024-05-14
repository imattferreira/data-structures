package structures

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

	for i < q.size {
		q.items[i] = nil
	}

	q.front = nil
	q.tail = nil
	q.size = 0
}

func (q *Queue) dequeue() (*int64, error) {
	if q.empty() {
		return nil, errors.New("[queue] is empty!")
	}

	var item = q.items[q.size]
	q.size--
	q.items[q.size] = nil
	q.tail = q.items[q.size-1]

	return item, nil
}

func (q *Queue) enqueue(item *int64) (*int64, error) {
	if q.full() {
		return nil, errors.New("[queue] is full!")
	}

	q.items[q.size] = item
	q.size++
	q.tail = item

	return item, nil
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
	}
}
