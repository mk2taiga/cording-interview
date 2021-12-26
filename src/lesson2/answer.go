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

// 2.4
func partition(node *Node, x int) *Node {
	if node == nil {
		return nil
	}

	head := node
	tail := node
	for node != nil {
		next := node.next

		if node.data < x {
			// 新しいデータをheadにする。
			node.next = head
			head = node
		} else {
			// 現状の末尾にnodeをつなげる
			tail.next = node
			// 新しい末尾
			tail = node
		}
		node = next
	}
	tail.next = nil
	return head
}

// 2.5
func addList(node1, node2 *Node, carry int) *Node {
	if node1 == nil && node2 == nil && carry == 0 {
		return nil
	}

	result := &Node{}

	value := carry
	if node1 != nil {
		value += node1.data
	}
	if node2 != nil {
		value += node2.data
	}

	result.data = value % 10
	if node1 != nil || node2 != nil {
		var n1 *Node
		var n2 *Node
		var nc int

		if node1 != nil {
			n1 = node1.next
		}
		if node2 != nil {
			n2 = node2.next
		}
		if value >= 10 {
			nc = 1
		}
		result.next = addList(n1, n2, nc)
	}

	return result
}

// 2.6
// TODO:
// 	逆順に並び替える。
// 	逆順に並び替えた要素と元々の要素が全て同じかどうかをチェックする。

// 2.7
// TODO:
// 	1. 二つの連結リストの長さを調べる。
//	2. 長い方の長さ - 短い方の長さ。
//	3. 長い方のリストを2の結果分進める。
// 	4. 長い方のリストと短い方のリストが同じ値になるまでループする。
//		for (short != long) {
//			short = short.next
//			long = long.next
//		}
// 	5. shortとlongが同じになっているかチェックする。同じならどちらかのnodeを返却して、そうでなければnilを返す。

// 2.8
// TODO:
//	1. ハッシュテーブルを用意する
//	2. map[node]bool として、すでに同じnodeがあればそれが、重複である。
// Other: 別の方法でfastランナーとslowランナーを使う方法もある。
