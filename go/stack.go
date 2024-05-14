package structures

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []*int64
	capacity,
	lastIndex,
	size uint64
}

func NewStack(capacity uint64) *Stack {
	return &Stack{
		capacity:  capacity,
		items:     []*int64{},
		lastIndex: 0,
		size:      0,
	}
}

func (s Stack) clear() {
	var i = s.size - 1

	for i != 0 {
		s.items[i] = nil
	}

	s.size = 0
	s.lastIndex = 0
}

func (s *Stack) empty() bool {
	return s.size == 0
}

func (s *Stack) full() bool {
	return s.size == s.capacity
}

func (s *Stack) insert(item *int64) (*int64, error) {
	if s.full() {
		return nil, errors.New("[stack] is full!")
	}

	s.lastIndex++
	s.items[s.lastIndex] = item

	return item, nil
}

func (s *Stack) pop() {}

func (s *Stack) show() {
	var i = s.size - 1

	for i != 0 {
		fmt.Println(s.items[i])
	}
}

func (s *Stack) top() *int64 {
	return s.items[s.lastIndex]
}
