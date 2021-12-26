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

// 2.2

func tail(n *Node, tailIdx int) *Node {
	var cnt int
	current := n
	for current != nil {
		cnt++
		current = current.next
	}

	idx := cnt - tailIdx
	if idx < 0 {
		return nil
	}

	current = n
	nowIdx := 0
	for current != nil {
		if idx == nowIdx {
			return current
		}
		nowIdx++
		current = current.next
	}

	return nil
}

// 2.3
func removeNode(node *Node) {
	if node == nil || node.next == nil {
		return
	}

	next := node.next
	node.data = next.data
	node.next = next.next
}
