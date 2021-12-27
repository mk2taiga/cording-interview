package lesson3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
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
}
