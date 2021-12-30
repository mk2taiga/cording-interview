package lesson10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	t.Run("test mergesort", func(t *testing.T) {
		list := []int{3, 1, 2}
		Mergesort(list)
		assert.Equal(t, []int{1, 2, 3}, list)
	})

	t.Run("test quicksort", func(t *testing.T) {
		list := []int{3, 1, 2, 0, 7}
		QuickSort(list, 0, 4)
		assert.Equal(t, []int{0, 1, 2, 3, 7}, list)
	})
}
