package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	items     []*int64
	capacity  uint64
	lastIndex int64
}

func NewStack(capacity uint64) *Stack {
	return &Stack{
		capacity:  capacity,
		items:     []*int64{},
		lastIndex: -1,
	}
}

func (s *Stack) clear() {
	for s.lastIndex >= 0 {
		s.items[s.lastIndex] = nil
		s.lastIndex--
	}
}

func (s *Stack) empty() bool {
	return s.lastIndex == -1
}

func (s *Stack) full() bool {
	return s.size() == int64(s.capacity)
}

func (s *Stack) insert(item *int64) error {
	if s.full() {
		return errors.New("[stack] is full")
	}

	s.lastIndex++
	s.items[s.lastIndex] = item

	return nil
}

func (s *Stack) pop() {
	if s.empty() {
		return
	}

	s.items[s.lastIndex] = nil
	s.lastIndex--
}

func (s *Stack) size() int64 {
	return s.lastIndex + 1
}

func (s *Stack) show() {
	i := s.lastIndex

	for i >= 0 {
		fmt.Println(s.items[i])
	}
}

func (s *Stack) top() *int64 {
	return s.items[s.lastIndex]
}
