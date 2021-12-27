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
}
