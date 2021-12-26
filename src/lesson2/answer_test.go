package lesson2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	t.Run("question 1", func(t *testing.T) {
		node := &Node{
			next: &Node{
				next: &Node{
					next: &Node{
						next: &Node{
							data: 1,
						},
						data: 4,
					},
					data: 3,
				},
				data: 2,
			},
			data: 1,
		}

		removeDuplicate(node)

		want := []int{1, 2, 3, 4}
		n := node
		i := 0
		for n != nil {
			assert.Equal(t, want[i], n.data)
			n = n.next
			i++
		}
	})

	t.Run("question 2", func(t *testing.T) {
		node := &Node{
			next: &Node{
				next: &Node{
					next: &Node{
						next: &Node{
							data: 1,
						},
						data: 4,
					},
					data: 3,
				},
				data: 2,
			},
			data: 1,
		}

		want := 4
		assert.Equal(t, want, tail(node, 2).data)
	})

	t.Run("question 3", func(t *testing.T) {
		deleteNode := &Node{
			next: &Node{
				next: &Node{data: 5},
				data: 4,
			},
			data: 3,
		}
		node := &Node{
			next: &Node{
				next: deleteNode,
				data: 2,
			},
			data: 1,
		}

		removeNode(deleteNode)
		want := []int{1, 2, 4, 5}
		n := node
		i := 0
		for n != nil {
			assert.Equal(t, want[i], n.data)
			n = n.next
			i++
		}
	})

	t.Run("question 4", func(t *testing.T) {
		node := &Node{
			next: &Node{
				next: &Node{
					next: &Node{
						next: &Node{
							next: &Node{
								next: &Node{
									data: 1,
								},
								data: 2,
							},
							data: 10,
						},
						data: 5,
					},
					data: 8,
				},
				data: 5,
			},
			data: 3,
		}

		got := partition(node, 5)
		want := []int{1, 2, 3, 5, 8, 5, 10}
		i := 0
		for got != nil {
			assert.Equal(t, want[i], got.data)
			got = got.next
			i++
		}
	})

	t.Run("question 5", func(t *testing.T) {
		node1 := &Node{
			next: &Node{
				next: &Node{data: 3},
				data: 2,
			},
			data: 1,
		}

		node2 := &Node{
			next: &Node{
				next: &Node{data: 9},
				data: 8,
			},
			data: 7,
		}

		got := addList(node1, node2, 0)
		want := []int{8, 0, 3, 1}
		i := 0
		for got != nil {
			assert.Equal(t, want[i], got.data)
			got = got.next
			i++
		}
	})
}
