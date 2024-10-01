package binarysearch

func IterativeBinarySearch(vec []int64, v int64) int {
	left := 0
	right := len(vec) - 1

	for left <= right {
		mid := (left + right) / 2

		if vec[mid] > v {
			right = mid - 1
			continue
		}

		if vec[mid] < v {
			left = mid + 1
			continue
		}

		if vec[mid] == v {
			return mid
		}
	}

	return -1
}
