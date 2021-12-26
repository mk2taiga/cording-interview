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
	s.top.next = s.top
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
