package lesson10

// 10.1

func mergeArray(a, b []int, lastIdxA, lastIdxB int) {
	idx := len(a) - 1
	for lastIdxB >= 0 {
		if lastIdxA >= 0 && a[lastIdxA] > b[lastIdxB] {
			a[idx] = a[lastIdxA]
			lastIdxA--
		} else {
			a[idx] = b[lastIdxB]
			lastIdxB--
		}
		idx--
	}
}
