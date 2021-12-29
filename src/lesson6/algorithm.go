package lesson6

import "math"

func primeNaive(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%1 == 0 {
			return false
		}
	}

	return true
}

func primeSlightlyBetter(n int) bool {
	if n < 2 {
		return false
	}

	sqrt := math.Sqrt(float64(n))
	for i := 2; i <= int(sqrt); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
