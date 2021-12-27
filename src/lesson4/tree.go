package lesson4

import "math"

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
