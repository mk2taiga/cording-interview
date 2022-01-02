package plactice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	t.Run("test quicksort", func(t *testing.T) {
		list := []int{3, 1, 2, 0, 7}
		quickSort(list, 0, 4)
		assert.Equal(t, []int{0, 1, 2, 3, 7}, list)

		list = []int{5, 1, 4, 3, 2}
		quickSort(list, 0, 4)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, list)

		list = []int{5, 1, 4, 3, 2, 10, 6, 102, 0, 100, 54}
		quickSort(list, 0, len(list)-1)
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 10, 54, 100, 102}, list)

		d := int64(33)
		prevSquare := d | (d >> 1)
		prevSquare |= prevSquare >> 2
		prevSquare |= prevSquare >> 4
		prevSquare |= prevSquare >> 8
		prevSquare |= prevSquare >> 16
		prevSquare |= prevSquare >> 32
		prevSquare -= prevSquare >> 1
		assert.Equal(t, int64(32), prevSquare)
	})
}
