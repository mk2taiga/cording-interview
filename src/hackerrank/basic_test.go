package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		assert.Equal(t, "00:05:45", timeConversion("12:05:45AM"))
	})

	t.Run("test 2", func(t *testing.T) {
		assert.Equal(t, 2147483648, flippingBits(9))
	})

	t.Run("test 3", func(t *testing.T) {
		assert.Equal(t, int32(2), birthday([]int32{1, 2, 1, 3, 2}, 3, 2))
	})

	t.Run("test 4", func(t *testing.T) {
		str := "baecdf"
		assert.Equal(t, "abcdef", stringSort(str))
	})
}
