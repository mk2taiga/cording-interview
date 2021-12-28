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
