// Big O = O**2
package sorting

func BubbleSort(v []int) []int {
	for i := range v {
		// small optimization: iterate until unsorted position
		for j := 0; j < len(v)-i-1; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
			}
		}
	}

	return v
}
