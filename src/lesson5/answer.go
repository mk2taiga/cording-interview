package lesson5

func getBit(num, i int) bool {
	return num&(1<<i) != 0
}

func setbit(num, i int) int {
	return num | 1<<i
}

func clearBit(num, i int) int {
	return num & ^(1 << i)
}

func clearHeadToI(num, i int) int {
	mask := (1 << i) - 1
	return num & mask
}

func crearItoTail(num, i int) int {
	mask := -1 << (i + 1)
	return num & mask
}

func updateBit(num, i int, val bool) int {
	var v int
	if val {
		v = 1
	}

	mask := ^(1 << i)
	return num&mask | v<<i
}

// 5.1

func updateBits(num, mum, j, i int) int {
	allOnes := ^0

	left := allOnes << (j + 1)
	right := (1 << i) - 1
	mask := left | right

	clean := num & mask
	shiftM := mum << i
	return clean | shiftM
}

// 5.2
// TODO: 10進数の少数を2進数に変換するロジックを書く
//		2 * num(少数)をして結果(re)がre>1ならば、2進数の結果に1を追加し、num = re-1をして再度2を掛けるを繰り返す。
//		最終的にnumが0より大きい間は繰り返す。

// 5.3
// TODO: 数字の1ビット目をシフトしながら確認していく。ループの限界値は、intのビット数とする。

// 5.5
// TODO: 2の累乗かどうかを表している。

// 5.6
func bitSwapRequired(a, b int) int {
	var cnt int
	for c := a ^ b; c != 0; c = c >> 1 {
		cnt += c & 1
	}
	return cnt
}
