package lesson4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	t.Run("test tree", func(t *testing.T) {
		node1 := &Node{val: 1}
		node2 := &Node{val: 2, parent: node1}
		node3 := &Node{val: 3, parent: node1}
		node1.left = node2
		node1.right = node3
		tree := &Tree{root: node1}

		tree.setDepth(tree.root, 0)
		tree.setHeight(tree.root)
		assert.Equal(t, 0, tree.root.depth)
		assert.Equal(t, 1, tree.root.height)
		assert.Nil(t, nil, tree.getSibling(tree.root))
		assert.Equal(t, 1, tree.root.right.depth)
		assert.Equal(t, 0, tree.root.left.height)
		assert.Equal(t, tree.root.right, tree.getSibling(tree.root.left))
		tree.preParse(tree.root)
		tree.inParse(tree.root)
		tree.postParse(tree.root)
	})

	t.Run("test heap", func(t *testing.T) {
		heap := Heap{
			list: []int{0, 4, 1, 3, 2, 16, 9, 10, 14, 8, 7},
			size: 10,
		}

		heap.buildMaxHeap()
		assert.Equal(t, []int{0, 16, 14, 10, 8, 7, 9, 3, 2, 4, 1}, heap.list)

		heap.insert(20)
		assert.Equal(t, []int{0, 20, 16, 10, 8, 14, 9, 3, 2, 4, 1, 7}, heap.list)

	})
}
