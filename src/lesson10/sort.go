package lesson10

func Mergesort(array []int) {
	helper := make([]int, len(array))
	mergesort(array, helper, 0, len(array)-1)
}

func mergesort(array, helper []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		mergesort(array, helper, low, mid)
		mergesort(array, helper, mid+1, high)
		merge(array, helper, low, mid, high)
	}
}

func merge(array, helper []int, low, mid, high int) {
	for i := low; i <= high; i++ {
		helper[i] = array[i]
	}

	left := low
	right := mid + 1
	current := low

	for left <= mid && right <= high {
		if helper[left] <= helper[right] {
			array[current] = helper[left]
			left++
		} else {
			array[current] = helper[right]
			right++
		}
		current++
	}

	remaining := mid - left
	for i := 0; i <= remaining; i++ {
		array[current+i] = helper[left+i]
	}
}

func QuickSort(arr []int, left, right int) {
	idx := partition(arr, left, right)
	if left < idx-1 {
		QuickSort(arr, left, idx-1)
	}
	if idx < right {
		QuickSort(arr, idx, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[(left+right)/2]
	for left <= right {
		for arr[left] < pivot {
			left++
		}

		for arr[right] > pivot {
			right--
		}

		if left <= right {
			tmp := arr[right]
			arr[right] = arr[left]
			arr[left] = tmp
			left++
			right--
		}
	}
	return left
}
