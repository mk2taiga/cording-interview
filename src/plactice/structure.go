package plactice

type Stack struct {
	data []int32
	size int
}

func NewStack() *Stack {
	return &Stack{data: make([]int32, 0, 4096), size: 0}
}

func (s *Stack) Push(n int32) {
	s.data = append(s.data, n)
	s.size++
}

func (s *Stack) Pop() bool {
	if s.IsEmpty() {
		return false
	}
	s.size--
	s.data = s.data[:s.size]
	return true
}

func (s *Stack) Top() int32 {
	return s.data[s.size-1]
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
