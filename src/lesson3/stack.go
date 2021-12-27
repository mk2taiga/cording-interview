package lesson3

type MyStack struct {
	top  *Node
	data interface{}
	size int
}

func NewMyStack(data interface{}) MyStack {
	return MyStack{data: data}
}

func (s *MyStack) pop() *Node {
	if s.top == nil {
		return nil
	}

	item := s.top
	s.top = s.top.next
	s.size--
	return item
}

func (s *MyStack) push(data interface{}) {
	node := &Node{data: data}
	node.next = s.top
	s.top = node
	s.size++
}

func (s *MyStack) peek() interface{} {
	if s.top == nil {
		return nil
	}

	return s.top.data
}

func (s *MyStack) isEmpty() bool {
	return s.top == nil
}

func (s *MyStack) sort() {
	r := NewMyStack(nil)
	for !s.isEmpty() {
		tmp := s.pop().data
		for !r.isEmpty() && tmp.(int) > r.peek().(int) {
			s.push(r.pop().data)
		}
		r.push(tmp)
	}

	for !r.isEmpty() {
		s.push(r.pop().data)
	}
}

type StackUseList struct {
	stack []interface{}
	cap   int
	top   int
}

func NewStackUseList(cap int) StackUseList {
	return StackUseList{
		stack: make([]interface{}, cap),
		cap:   cap,
	}
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

func (s *StackUseList) isFill() bool {
	return s.top == s.cap
}

func (s *StackUseList) isEmpty() bool {
	return s.top == 0
}
