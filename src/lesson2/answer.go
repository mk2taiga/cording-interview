package lesson2

type Node struct {
	next *Node
	data int
}

func (n *Node) append(d int) {
	end := &Node{
		data: d,
	}

	now := n
	for now.next != nil {
		now = now.next
	}
	now.next = end
}

func remove(head *Node, d int) *Node {
	n := head

	if n.data == d {
		return n.next
	}

	for n.next != nil {
		if n.next.data == d {
			n.next = n.next.next
			return head
		}
		n = n.next
	}

	return head
}

// 2.1
func removeDuplicate(n *Node) {
	dupMap := map[int]bool{}
	var before *Node
	for n != nil {
		if dupMap[n.data] {
			before.next = n.next
		} else {
			before = n
			dupMap[n.data] = true
		}

		n = n.next
	}
}
