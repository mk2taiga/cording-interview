package lesson3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	t.Run("question 1", func(t *testing.T) {
		stack := NewFixedMultiStack(3)
		assert.True(t, stack.push(0, 0))
		assert.True(t, stack.push(0, 1))
		assert.True(t, stack.push(0, 2))
		assert.True(t, stack.push(1, 3))
		assert.True(t, stack.push(1, 4))
		assert.True(t, stack.push(1, 5))
		assert.True(t, stack.push(2, 6))
		assert.True(t, stack.push(2, 7))
		assert.True(t, stack.push(2, 8))
		assert.False(t, stack.push(2, 9))

		val, err := stack.peek(0)
		assert.NoError(t, err)
		assert.Equal(t, 0, val)
		val, err = stack.peek(1)
		assert.NoError(t, err)
		assert.Equal(t, 3, val)
		val, err = stack.peek(2)
		assert.NoError(t, err)
		assert.Equal(t, 6, val)

		for i := 2; i >= 0; i-- {
			val, err = stack.pop(0)
			assert.NoError(t, err)
			assert.Equal(t, i, val)
			val, err = stack.pop(1)
			assert.NoError(t, err)
			assert.Equal(t, i+3, val)
			val, err = stack.pop(2)
			assert.NoError(t, err)
			assert.Equal(t, i+6, val)
		}
	})

	t.Run("question 2", func(t *testing.T) {
		stack := NewStackWithMin()
		stack.push(3)
		stack.push(2)
		stack.push(1)

		assert.Equal(t, 1, stack.min())
		assert.Equal(t, 1, stack.pop())
		assert.Equal(t, 2, stack.min())
		assert.Equal(t, 2, stack.pop())
		assert.Equal(t, 3, stack.min())
		assert.Equal(t, 3, stack.pop())
	})

	t.Run("question 4", func(t *testing.T) {
		q := NewMyQueueByStack()
		q.add(1)
		q.add(2)
		q.add(3)

		hoge := q.remove().(int)
		assert.Equal(t, 1, hoge)
		assert.Equal(t, 2, q.remove().(int))
		assert.Equal(t, 3, q.remove().(int))
	})
}
