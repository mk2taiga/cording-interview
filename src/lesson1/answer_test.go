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

	t.Run("question 5", func(t *testing.T) {
		str1 := "pile"
		str2 := "pil"
		assert.True(t, isOneChange(str1, str2))

		str1 = "pile"
		str2 = "piles"
		assert.True(t, isOneChange(str1, str2))

		str1 = "pile"
		str2 = "bile"
		assert.True(t, isOneChange(str1, str2))

		str1 = "pale"
		str2 = "bake"
		assert.False(t, isOneChange(str1, str2))
	})

	t.Run("question 6", func(t *testing.T) {
		str := "aabcccccaaa"
		want := "a2b1c5a3"
		assert.Equal(t, want, strPress(str))

		str = "abcdefg"
		want = "abcdefg"
		assert.Equal(t, want, strPress(str))
	})

	t.Run("question 7", func(t *testing.T) {
		matrix := [][]int32{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}
		want := [][]int32{
			{13, 9, 5, 1},
			{14, 10, 6, 2},
			{15, 11, 7, 3},
			{16, 12, 8, 4},
		}
		rotationMatrix(matrix)
		assert.Equal(t, want, matrix)
	})

	t.Run("question 8", func(t *testing.T) {
		matrix := [][]int32{
			{1, 2, 3, 4},
			{5, 0, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 0, 16},
		}
		want := [][]int32{
			{1, 0, 0, 4},
			{0, 0, 0, 0},
			{9, 0, 0, 12},
			{0, 0, 0, 0},
		}
		toZero(matrix)
		assert.Equal(t, want, matrix)
	})

	t.Run("question 1.9", func(t *testing.T) {
		str1 := "waterbottle"
		str2 := "erbottlewat"
		assert.True(t, isRotation(str1, str2))

		str1 = "waterbottle"
		str2 = "erabottlewat"
		assert.False(t, isRotation(str1, str2))
	})
}
