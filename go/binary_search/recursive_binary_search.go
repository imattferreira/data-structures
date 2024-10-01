package binarysearch

func RecursiveBinarySearch(vec []int64, v int64, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if vec[mid] > v {
		return RecursiveBinarySearch(vec, v, left, mid-1)
	}

	if vec[mid] < v {
		return RecursiveBinarySearch(vec, v, mid+1, right)
	}

	if vec[mid] == v {
		return mid
	}

	return -1
}
