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
	n := s.stack.pop()
	if n.data == s.min() {
		s.minStack.pop()
	}

	return n.data.(int)
}

func (s *StackWithMin) min() int {
	if s.minStack.isEmpty() {
		return math.MaxInt
	}

	return s.minStack.peek().(int)
}

// 3.3
// TODO:
// 	Stack を管理するリストが必要。
//  * push するときに、最後のスタックを取得してきて、それがnilではなく、fillではなかったらそのまま追加、そうでなければ
//  新しいスタックを追加。
//  * popの時は、スタックが空になったら、そのスタックを削除する。
// 	* popAtでは、スタックのサイズが指定したidxよりも大きかったら、指定したidx以上のスタックから、データをおろしてこないといけない。
//  Nodeが連結リストなので、上のスタックから一つおろしてくれば問題ない。連結リストは、追加や削除が一定時間でできるので非常に便利。
//  * この実装では、Nodeは走行の連結リストになっていないといけない。

// 3.4

// MyQueueByStack キューはFIFO
type MyQueueByStack struct {
	// 新しい要素が上に来るスタック
	stackNew MyStack
	// 古い要素が上に来るスタック
	stackOld MyStack
}

func NewMyQueueByStack() MyQueueByStack {
	return MyQueueByStack{
		stackNew: NewMyStack(nil),
		stackOld: NewMyStack(nil),
	}
}

func (q *MyQueueByStack) size() int {
	return q.stackNew.size + q.stackOld.size
}

func (q *MyQueueByStack) add(data interface{}) {
	q.stackNew.push(data)
}

func (q *MyQueueByStack) shiftStacks() {
	if q.stackOld.isEmpty() {
		for !q.stackNew.isEmpty() {
			q.stackOld.push(q.stackNew.pop().data)
		}
	}
}

func (q *MyQueueByStack) peek(data interface{}) interface{} {
	q.shiftStacks()
	return q.stackOld.peek()
}

func (q *MyQueueByStack) remove() interface{} {
	q.shiftStacks()
	return q.stackOld.pop().data
}

// 3.6
// TODO:
//		1. 犬と猫のみをキューで管理できればいいので、AnimalQueue のなかに二つの二つのリストを持つ。
//		2. 動物のデータは、タイムスタンプを持つ
//  	3. pushAny はそれぞれのリストから0番目を取得して、タイムスタンプで比較して古い方だけを取り出す。
