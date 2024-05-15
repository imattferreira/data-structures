package circularqueue

import (
	"errors"
	"fmt"
)

type CircularQueue struct {
	capacity,
	frontIndex,
	tailIndex,
	size uint64
	items []*int64
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
	i := cq.frontIndex

	for i != cq.tailIndex {
		cq.items[i] = nil
		i = cq.nextIndex(i)
	}
}

func (cq *CircularQueue) dequeue() {
	if cq.empty() {
		return
	}

	cq.frontIndex = cq.nextIndex(cq.frontIndex)
	cq.size--
}

func (cq *CircularQueue) enqueue(value *int64) error {
	if cq.full() {
		return errors.New("[circular-queue] is full")
	}

	cq.tailIndex = cq.nextIndex(cq.tailIndex)
	cq.items[cq.tailIndex] = value
	cq.size++

	return nil
}

func (cq *CircularQueue) front() *int64 {
	return cq.items[cq.frontIndex]
}

func (cq *CircularQueue) tail() *int64 {
	return cq.items[cq.tailIndex]
}

func (cq *CircularQueue) show() {
	i := cq.frontIndex

	for i != cq.tailIndex {
		fmt.Println(cq.items[i])
		i = cq.nextIndex(i)
	}
}

func (cq *CircularQueue) nextIndex(index uint64) uint64 {
	if index == cq.capacity-1 {
		return 0
	}

	return index + 1
}

func (cq *CircularQueue) full() bool {
	return cq.size == cq.capacity
}

func (cq *CircularQueue) empty() bool {
	return cq.size == 0
}
