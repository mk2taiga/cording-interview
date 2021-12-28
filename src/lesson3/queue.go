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

func (q *MyQueue) Add(data interface{}) {
	node := &Node{data: data}
	if q.last != nil {
		q.last.next = node
	}
	q.last = node

	if q.first == nil {
		q.first = node
	}
}

func (q *MyQueue) Remove() interface{} {
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

func (q *MyQueue) Peek() interface{} {
	if q.first == nil {
		return nil
	}

	return q.first.data
}

func (q *MyQueue) IsEmpty() bool {
	return q.first == nil
}

type MyQueueUseList struct {
	head  int
	tail  int
	max   int
	queue []interface{}
}

func NewMyQueueUseList(max int) MyQueueUseList {
	return MyQueueUseList{
		max:   max,
		queue: make([]interface{}, max),
	}
}

func (q *MyQueueUseList) isFull() bool {
	return q.head == (q.tail+1)%q.max
}

func (q *MyQueueUseList) add(data interface{}) {
	if q.isFull() {
		return
	}

	q.queue[q.tail] = data
	q.tail++
	if q.tail == q.max {
		q.tail = 0
	}
}

func (q *MyQueueUseList) remove() interface{} {
	if q.isEmpty() {
		return nil
	}

	data := q.queue[q.head]
	q.head++
	if q.head == q.max {
		q.head = 0
	}
	return data
}

func (q *MyQueueUseList) peek() interface{} {
	if q.isEmpty() {
		return nil
	}

	return q.queue[q.head]
}

func (q *MyQueueUseList) isEmpty() bool {
	return q.head == q.tail
}
