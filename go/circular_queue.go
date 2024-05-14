package structures

import (
	"errors"
	"fmt"
)

type CircularQueue struct {
	capacity, frontIndex, tailIndex, size uint64
	items                                 []*int64
}

func NewCircularQueue(capacity uint64) *CircularQueue {
	return &CircularQueue{
		capacity:   capacity,
		frontIndex: 0,
		tailIndex:  0,
		size:       0,
		items:      []*int64{},
	}
}

func (cq *CircularQueue) clear() {
	var i = cq.frontIndex

	for i != cq.tailIndex {
		cq.items[i] = nil

		i = (func() uint64 {
			if cq.isLast(i) {
				return 0
			}

			return i + 1
		})()
	}
}

func (cq *CircularQueue) dequeue() (*int64, error) {
	if cq.empty() {
		return nil, errors.New("[circular-queue] is empty!")
	}

	var item = cq.items[cq.frontIndex]

	cq.frontIndex = (func() uint64 {
		if cq.isLast(cq.frontIndex) {
			return 0
		}

		return cq.frontIndex + 1
	})()
	cq.size--

	return item, nil
}

func (cq *CircularQueue) enqueue(value *int64) (*int64, error) {
	if cq.full() {
		return nil, errors.New("[circular-queue] is full!")
	}

	cq.tailIndex = (func() uint64 {
		if cq.isLast(cq.tailIndex) {
			return 0
		}

		return cq.tailIndex - 1
	})()
	cq.items[cq.tailIndex] = value

	return value, nil
}

func (cq *CircularQueue) front() *int64 {
	return cq.items[cq.frontIndex]
}

func (cq *CircularQueue) tail() *int64 {
	return cq.items[cq.tailIndex]
}

func (cq *CircularQueue) show() {
	var i = cq.frontIndex

	for i != cq.tailIndex {
		fmt.Println(cq.items[i])

		i = (func() uint64 {
			if cq.isLast(i) {
				return 0
			}

			return i + 1
		})()
	}
}

func (cq *CircularQueue) isLast(index uint64) bool {
	return index == cq.capacity-1
}

func (cq *CircularQueue) full() bool {
	return cq.size == cq.capacity
}

func (cq *CircularQueue) empty() bool {
	return cq.size == 0
}
