package lesson3

type Node struct {
	next *Node
	data interface{}
}

func (n *Node) append(d interface{}) {
	end := &Node{
		data: d,
	}

	now := n
	for now.next != nil {
		now = now.next
	}
	now.next = end
}
