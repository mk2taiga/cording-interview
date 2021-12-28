package lesson4

import (
	"cording-interview/src/lesson3"
	"math"
)

type Node struct {
	parent, left, right *Node
	val                 int
	depth               int
	height              int
}

type Tree struct {
	root *Node
}

func (t Tree) setDepth(u *Node, d int) {
	if u == nil {
		return
	}

	u.depth = d
	t.setDepth(u.left, d+1)
	t.setDepth(u.right, d+1)
}

func (t Tree) setHeight(u *Node) int {
	h1, h2 := 0, 0
	if u.left != nil {
		h1 = t.setHeight(u.left) + 1
	}
	if u.right != nil {
		h2 = t.setHeight(u.right) + 1
	}

	u.height = int(math.Max(float64(h1), float64(h2)))
	return u.height
}

func (t Tree) getSibling(u *Node) *Node {
	if u.parent == nil {
		return nil
	}

	if u.parent.left != nil && u.parent.left != u {
		return u.parent.left
	}

	if u.parent.right != nil && u.parent.right != u {
		return u.parent.right
	}

	return nil
}

func (t Tree) preParse(u *Node) {
	if u == nil {
		return
	}

	println(u.val)
	t.preParse(u.left)
	t.preParse(u.right)
}

func (t Tree) inParse(u *Node) {
	if u == nil {
		return
	}

	t.preParse(u.left)
	println(u.val)
	t.preParse(u.right)
}

func (t Tree) postParse(u *Node) {
	if u == nil {
		return
	}

	t.preParse(u.left)
	t.preParse(u.right)
	println(u.val)
}

type BinTree struct {
	root *Node
}

func (t BinTree) insert(val int) {
	newNode := &Node{val: val}

	var leaf *Node
	child := t.root
	for child != nil {
		leaf = child
		if newNode.val < child.val {
			child = child.left
		} else {
			child = child.right
		}
	}

	newNode.parent = leaf
	if leaf == nil {
		t.root = newNode
	} else {
		if newNode.val < leaf.val {
			leaf.right = newNode
		} else {
			leaf.left = newNode
		}
	}
}

func (t BinTree) find(val int) *Node {
	u := t.root
	for u != nil && u.val != val {
		if val < u.val {
			u = u.left
		} else {
			u = u.right
		}
	}

	return u
}

func (t BinTree) treeMin(u *Node) *Node {
	for u.left != nil {
		u = u.left
	}
	return u
}

func (t BinTree) treeSuccessor(u *Node) *Node {
	if u.right != nil {
		return t.treeMin(u)
	}

	p := u.parent
	for p != nil && p.right == u {
		u = p
		p = p.parent
	}

	return p
}

func (t BinTree) delete(u *Node) {
	var target *Node
	var targetChild *Node

	if u.left == nil || u.right == nil {
		target = u
	} else {
		target = t.treeSuccessor(u)
	}

	if target.left != nil {
		targetChild = target.left
	} else {
		targetChild = target.right
	}

	if targetChild != nil {
		targetChild.parent = target.parent
	}

	if target.parent == nil {
		t.root = targetChild
	} else {
		if target == target.parent.left {
			target.parent.left = targetChild
		} else {
			target.parent.right = targetChild
		}
	}

	if target != u {
		u.val = target.val
	}
}

type Heap []int

func (h Heap) maxHeapify(idx int) {
	l := 2 * idx
	r := 2*idx + 1

	largest := idx
	if l <= len(h) && h[l] > h[idx] {
		largest = l
	}
	if r <= len(h) && h[r] > h[idx] {
		largest = r
	}

	if largest != idx {
		h[idx], h[largest] = h[largest], h[idx]
		h.maxHeapify(largest)
	}
}

// TODO: 優先度付きキューの実装(insert)
//	キューの最後に値を追加する
//	キューのparent(i/2)と比較して、新しい値の方が大きければ入れ替える。
//	parentのindexが1より大きい間繰り返す。

// TODO: 優先度付きキューの実装(最大値削除)
// 	0インデックスを取得して、一番小さい値を0idxに格納する。
//  0番目からひたすらmaxHeapify(0)を実行する。

// Graph is a graph with n vertices and edges.
type Graph struct {
	n     int
	edges [][]int
	color []int
}

func (g Graph) dfs() {
	for i := 0; i < g.n; i++ {
		g.color[i] = 0
	}

	for i := 0; i < g.n; i++ {
		if g.color[i] == 0 {
			g.dfsVisit(i)
		}
	}
}

func (g Graph) dfsVisit(pos int) {
	g.color[pos] = 1
	for visit := 0; visit < g.n; visit++ {
		if g.edges[pos][visit] == 0 {
			continue
		}
		if g.color[visit] == 0 {
			g.dfsVisit(visit)
		}
	}

	g.color[pos] = 2
}

func (g Graph) bfs(start int) {
	q := lesson3.NewMyQueue(nil)
	q.Add(start)
	for i := 0; i < g.n; i++ {
		g.color[i] = 0
	}
	g.color[start] = 1

	var u int
	for !q.IsEmpty() {
		u = q.Remove().(int)
		for v := 0; v < g.n; v++ {
			if g.edges[u][v] == 0 {
				continue
			}
			if g.color[v] != 0 {
				continue
			}
			g.color[v] = 1
			q.Add(v)
		}
		g.color[u] = 2
	}
}
