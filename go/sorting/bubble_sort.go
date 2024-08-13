package sorting

func BubbleSort(v []int) []int {
	swapped := false

	for i := range v {
		/**
		 * small optimization: iterate only over i until last index instead
		 * of all elems
		 */
		for j := 0; j < i-1; j++ {
			if v[j] > v[j+1] {
				tmp := v[j]
				v[j] = v[j+1]
				v[j+1] = tmp
			}
		}

		// small optimization: consider slice as already sorted
		if !swapped {
			break
		}
	}

	return v
}
