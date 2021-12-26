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

	t.Run("question 2", func(t *testing.T) {
		str1 := "bacdefg"
		str2 := "bacdgfe"
		assert.True(t, isSortedStr(str1, str2))

		str1 = "bacdefg"
		str2 = "bacggfe"
		assert.False(t, isSortedStr(str1, str2))
	})

	t.Run("question 2(other answer)", func(t *testing.T) {
		str1 := "bacdefg"
		str2 := "bacdgfe"
		assert.True(t, isSortedStrByASCII(str1, str2))

		str1 = "bacdefg"
		str2 = "bacggfe"
		assert.False(t, isSortedStrByASCII(str1, str2))
	})

	t.Run("question 3", func(t *testing.T) {
		str := "My John Smith "

		assert.Equal(t, "My%20John%20Smith%20", spaceToCode(str))
	})

	t.Run("question 4", func(t *testing.T) {
		str := "Tact Coa"
		assert.True(t, isReplyString(str))

		str = "Tact Coaa"
		assert.False(t, isReplyString(str))
	})
}
