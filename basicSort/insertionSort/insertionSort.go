package insertionSort

func insertionSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	for i := 1; i < len(a); i++ {
		current := i
		for j := i - 1; j >= 0; j-- {
			if a[j] >= a[current] {
				a[j], a[current] = a[current], a[j]
				current = j
			} else {
				break
			}
		}
	}

	return a
}
