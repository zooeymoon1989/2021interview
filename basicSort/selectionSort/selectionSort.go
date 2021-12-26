package selectionSort

func selectionSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	for i := len(a) - 1; i > 0; i-- {
		max := i
		for j := 0; j < i-1; j++ {
			if a[max] < a[j] {
				max = j
			}
		}
		if max != i {
			a[max], a[i] = a[i], a[max]
		}
	}

	return a
}
