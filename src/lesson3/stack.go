package lesson3

import "errors"

type MyStack struct {
	top  *Node
	data interface{}
}

func NewMyStack(data interface{}) MyStack {
	return MyStack{data: data}
}

func (s *MyStack) pop() (*Node, error) {
	if s.top == nil {
		return nil, errors.New("top not found")
	}

	item := s.top
	s.top = s.top.next
	return item, nil
}

func (s *MyStack) push(data interface{}) {
	node := &Node{data: data}
	node.next = s.top
	s.top = node
}

func (s *MyStack) peek() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("error occurs")
	}

	return s.top.data, nil
}

func (s *MyStack) isEmpty() bool {
	return s.top == nil
}

type StackUseList struct {
	stack []interface{}
	top   int
}

func NewStackUseList(cap int) StackUseList {
	return StackUseList{stack: make([]interface{}, cap)}
}

func (s *StackUseList) push(data interface{}) {
	s.stack[s.top] = data
	s.top++
}

func (s *StackUseList) pop() interface{} {
	if s.isEmpty() {
		return nil
	}

	s.top--
	data := s.stack[s.top]
	s.stack[s.top] = nil
	return data
}

func (s *StackUseList) peek() interface{} {
	if s.isEmpty() {
		return nil
	}

	return s.stack[s.top]
}

func (s *StackUseList) isEmpty() bool {
	return s.top == 0
}
