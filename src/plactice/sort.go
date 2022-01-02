package plactice

func quickSort(arr []int, left, right int) {
	idx := partition(arr, left, right)
	if left < idx-1 {
		quickSort(arr, left, idx-1)
	}

	// Why don't we add one to idx?
	// This is because we can't know where the pivot is set in list.
	if right > idx {
		quickSort(arr, idx, right)
	}
}

func partition(arr []int, left, right int) int {
	// left=0, right=5, div = 2
	// pivot is the value at 2nd index of the arr.
	pivot := arr[(left+right)/2]

	// e.g.
	//	loop=1: [5, 1, 4, 3, 2], left=0, right=5
	//	loop=2: [2, 1, 4, 3, 5], left=1, right=4
	//	loop=3: [2, 1, 3, 4, 5], left=3, right=2, so loop will be finished
	for left <= right {
		// e.g.
		// 	left=0
		// 	left=2
		for arr[left] < pivot {
			left++
		}

		// e.g.
		// 	right=4
		//	right=3
		for arr[right] > pivot {
			right--
		}

		// If the left is over the right, we are not necessary to exchange each values.
		if left <= right {
			tmp := arr[left]
			arr[left] = arr[right]
			arr[right] = tmp
			// Now value is already checked, so we proceed the left and right index.
			left++
			right--
		}
	}

	return left
}
