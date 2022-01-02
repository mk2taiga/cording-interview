package hackerrank

import (
	"fmt"
	"sort"
	"strconv"
)

func timeConversion(s string) string {
	// Write your code here
	timeStr := s[:len(s)-2]
	hour, _ := strconv.Atoi(timeStr[:2])

	if s[len(s)-2:] == "AM" {
		// 12の時は12時間マイナスする。
		if hour == 12 {
			hour -= 12
		}
	} else {
		hour += 12
	}

	res := fmt.Sprintf("%02d:%s", hour, timeStr[3:])
	return res
}

func flippingBits(n int64) int64 {
	// Write your code here
	// 00000001 -> 11111110
	reverseN := ^uint32(n)
	// 0*32+1*32
	allOne := ^uint32(0)
	res := int64(reverseN & allOne)
	return res
}

func int32Sort(arr []int32) {
	array := []int32{1, 4, 2, 9, 3, 8, 7, 6, 10, 5}
	fmt.Println(array)
	sort.Slice(array, func(i, j int) bool { return array[i] < array[j] })
	fmt.Println(array)
}

func stringSort(str string) string {
	v := []rune(str)
	sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
	return string(v)
}

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	var cnt int32
	first := int32(0)
	last := first + m
	for ; last <= int32(len(s)); last++ {
		sum := int32(0)
		list := s[first:last]
		for _, v := range list {
			sum += v
		}

		if sum == d {
			cnt++
		}
		first++
	}

	return cnt
}
