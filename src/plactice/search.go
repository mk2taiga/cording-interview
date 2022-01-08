package plactice

func BinarySearch(key int, nums []int) bool {
	var left, right = 0, len(nums)

	for left < right {
		var mid = (left + right) / 2
		if nums[mid] == key {
			return true
		} else if key < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
}
