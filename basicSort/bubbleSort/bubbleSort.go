package bubbleSort

func bubbleSort(a []int) []int {

	if len(a) < 2 {
		return a
	}

	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}
