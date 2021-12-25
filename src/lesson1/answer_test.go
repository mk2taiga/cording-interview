package lesson1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	t.Run("question 1", func(t *testing.T) {
		testStr := "abcdaebfg"
		assert.False(t, isNotDuplicate(testStr))

		testStr = "abcdefghij"
		assert.True(t, isNotDuplicate(testStr))
	})

	t.Run("question 1(other answer)", func(t *testing.T) {
		testStr := "abcdaebfg"
		assert.False(t, isNotDuplicateByByte(testStr))

		testStr = "abcdefghij"
		assert.True(t, isNotDuplicateByByte(testStr))
	})
}
