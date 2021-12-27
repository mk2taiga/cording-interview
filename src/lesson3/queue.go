package lesson3

type MyQueue struct {
	top   *Node
	first *Node
	last  *Node
	data  interface{}
}

func NewMyQueue(data interface{}) MyQueue {
	return MyQueue{data: data}
}

func (q *MyQueue) add(data interface{}) {
	node := &Node{data: data}
	if q.last != nil {
		q.last.next = node
	}
	q.last = node

	if q.first == nil {
		q.first = node
	}
}

func (q *MyQueue) remove() interface{} {
	if q.first == nil {
		return nil
	}

	data := q.first.data
	q.first = q.first.next
	if q.first == nil {
		q.last = nil
	}
	return data
}

func (q *MyQueue) peek() interface{} {
	if q.first == nil {
		return nil
	}

	return q.first.data
}

func (q *MyQueue) isEmpty() bool {
	return q.first == nil
}
