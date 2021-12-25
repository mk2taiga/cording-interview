package lesson0

type ArrayList struct {
	cnt    int
	buffer []int
}

func NewArrayList(cap *int) *ArrayList {
	if cap == nil {
		return &ArrayList{buffer: make([]int, 4)}
	} else {
		return &ArrayList{buffer: make([]int, *cap)}
	}
}

func (l *ArrayList) Cap() int {
	return len(l.buffer)
}

func (l *ArrayList) Add(value int) bool {
	if l.Cap() == 0 {
		l.buffer = make([]int, 4)
	} else if l.Cap() == l.cnt {
		src := l.buffer
		l.buffer = make([]int, l.Cap()*2)
		copy(l.buffer, src)
	}

	l.buffer[l.cnt] = value
	l.cnt++
	return true
}

func (l *ArrayList) Dump() []int {
	return l.buffer
}
