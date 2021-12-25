package lesson1

// 1.1
func isNotDuplicate(value string) bool {
	// 文字はASCIIとする。
	if len(value) > 128 {
		return false
	}

	charSet := make([]bool, 128)
	for _, v := range value {
		if b := charSet[v]; b == true {
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
