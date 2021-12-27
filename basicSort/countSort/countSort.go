package countSort

func countSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	min, max := a[0], a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}

		if a[i] > max {
			max = a[i]
		}
	}

	return a
}
