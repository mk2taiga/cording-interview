package lesson3

import (
	"errors"
	"math"
)

// 3.1
type FixedMultiStack struct {
	numberOfStack int
	stackCap      int
	values        []int
	sizes         []int
}

func NewFixedMultiStack(cap int) FixedMultiStack {
	return FixedMultiStack{
		numberOfStack: 3,
		stackCap:      cap,
		values:        make([]int, 3*cap),
		sizes:         make([]int, 3),
	}
}

func (s *FixedMultiStack) push(num int, data int) bool {
	if s.isFull(num) {
		return false
	}

	s.sizes[num]++
	s.values[s.idxOfTop(num)] = data
	return true
}

func (s *FixedMultiStack) pop(num int) (int, error) {
	if s.isEmpty(num) {
		return 0, errors.New("nothing")
	}

	idx := s.idxOfTop(num)
	val := s.values[idx]
	s.values[idx] = 0
	s.sizes[num]--
	return val, nil
}

func (s *FixedMultiStack) peek(num int) (int, error) {
	if s.isEmpty(num) {
		return 0, errors.New("nothing")
	}

	return s.values[s.idxOfTop(num)], nil
}

func (s *FixedMultiStack) isEmpty(num int) bool {
	return s.sizes[num] == 0
}

func (s *FixedMultiStack) isFull(num int) bool {
	return s.sizes[num] == s.stackCap
}

func (s *FixedMultiStack) idxOfTop(num int) int {
	return s.sizes[num] + s.stackCap*num - 1
}

// 3.2
type StackWithMin struct {
	stack    MyStack
	minStack MyStack
}

func NewStackWithMin() StackWithMin {
	return StackWithMin{
		stack:    NewMyStack(nil),
		minStack: NewMyStack(nil),
	}
}

func (s *StackWithMin) push(data int) {
	s.stack.push(data)
	if data <= s.min() {
		s.minStack.push(data)
	}
}

func (s *StackWithMin) pop() int {
	n, _ := s.stack.pop()
	if n.data == s.min() {
		s.minStack.pop()
	}

	return n.data.(int)
}

func (s *StackWithMin) min() int {
	if s.minStack.isEmpty() {
		return math.MaxInt
	}

	val, _ := s.minStack.peek()
	return val.(int)
}
