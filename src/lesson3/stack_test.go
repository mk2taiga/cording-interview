package lesson3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("question 5", func(t *testing.T) {
		s := NewMyStack(nil)
		s.push(5)
		s.push(4)
		s.push(2)
		s.push(1)
		s.push(3)

		s.sort()
		assert.Equal(t, s.pop().data, 5)
		assert.Equal(t, s.pop().data, 4)
		assert.Equal(t, s.pop().data, 3)
		assert.Equal(t, s.pop().data, 2)
		assert.Equal(t, s.pop().data, 1)
	})
}
