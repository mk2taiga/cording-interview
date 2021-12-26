package lesson1

import (
	"sort"
	"strconv"
	"strings"
)

// 1.1
func isNotDuplicate(value string) bool {
	// 文字はASCIIとする。
	if len(value) > 128 {
		return false
	}

	charSet := make([]bool, 128)
	for _, v := range value {
		if charSet[v] {
			return false
		}
		charSet[v] = true
	}

	return true
}

func isNotDuplicateByByte(value string) bool {
	checker := 0
	for i := 0; i < len(value); i++ {
		c := value[i] - "a"[0]
		if checker&(1<<c) > 0 {
			return false
		}

		checker |= 1 << c
	}

	return true
}

// 1.2

type SortRune []rune

func (r SortRune) Len() int           { return len(r) }
func (r SortRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r SortRune) Less(i, j int) bool { return r[i] < r[j] }

func isSortedStr(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	sorted1 := []rune(str1)
	sorted2 := []rune(str2)
	sort.Sort(SortRune(sorted1))
	sort.Sort(SortRune(sorted2))

	return string(sorted1) == string(sorted2)
}

func isSortedStrByASCII(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	charSet := make([]int, 128)
	for _, s := range str1 {
		charSet[s]++
	}

	for _, s := range str2 {
		charSet[s]--
		if charSet[s] < 0 {
			return false
		}
	}

	return true
}

// 1.3

func spaceToCode(str string) string {
	var newStr string
	for _, s := range str {
		if string(s) == " " {
			newStr += "%20"
		} else {
			newStr += string(s)
		}
	}

	return newStr
}

// 1.4
func isReplyString(str string) bool {
	lowerStr := strings.ToLower(str)
	strMap := map[rune]int{}
	strLen := 0
	for _, r := range lowerStr {
		if string(r) == " " {
			continue
		}

		strLen++
		strMap[r]++
	}

	foundOdd := false
	for _, v := range strMap {
		if v%2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}

	return true
}

// 1.5

func isOneChange(str1, str2 string) bool {
	var bitVector int

	lowerStr1 := strings.ToLower(str1)
	lowerStr2 := strings.ToLower(str2)
	minChar := int32("a"[0])
	maxChar := int32("z"[0])
	for _, v := range lowerStr1 {
		if v < minChar && v > maxChar {
			continue
		}
		bitVector |= 1 << (v - minChar)
	}

	for _, v := range lowerStr2 {
		if v < minChar && v > maxChar {
			continue
		}
		bitVector &= ^(1 << (v - minChar))
	}

	return (bitVector & (bitVector - 1)) == 0
}

// 1.6

// StringBuilder を使って処理を文字列結合に使う時間を最適化
func strPress(str string) string {
	var strCnt int
	builder := strings.Builder{}

	for i, v := range str {
		strCnt++

		if i+1 >= len(str) || v != int32(str[i+1]) {
			builder.WriteString(string(v))
			builder.WriteString(strconv.Itoa(strCnt))
			strCnt = 0
		}
	}

	result := builder.String()
	if len(result) >= len(str) {
		return str
	}

	return result
}

// 1.7
func rotationMatrix(matrix [][]int32) bool {
	// 行列の形式が正しくなかったら false
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return false
	}

	n := len(matrix)
	// 行列の層の数は辺の半分までのはず。
	for layer := 0; layer < n/2; layer++ {
		// 層の最初の位置
		first := layer
		// 層の最後の位置
		last := n - 1 - layer
		for i := first; i < last; i++ {
			// i の初期値は first なので引いてあげないと正式なoffsetがわからない。
			offset := i - first

			// 左上
			top := matrix[first][i]
			// 左上に左下
			matrix[first][i] = matrix[last-offset][first]
			// 左下に右下
			matrix[last-offset][first] = matrix[last][last-offset]
			// 右下に右上
			matrix[last][last-offset] = matrix[i][last]
			// 右上に左上
			matrix[i][last] = top
		}
	}
	return true
}

// 1.8

func toZero(matrix [][]int32) {
	rowMap := map[int]bool{}
	colMap := map[int]bool{}

	for i, row := range matrix {
		for j, col := range row {
			if col == 0 {
				rowMap[i] = true
				colMap[j] = true
			}
		}
	}

	for i, row := range matrix {
		for j := range row {
			if rowMap[i] || colMap[j] {
				matrix[i][j] = 0
			}
		}
	}
}
