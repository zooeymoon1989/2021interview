package shellSort

func shellSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	for i := len(a) / 2; i > 0; i /= 2 {
		for j := i; j < len(a); j++ {
			for k := j; k >= 0 && k-i >= 0; k -= i {
				if a[k] < a[k-i] {
					a[k], a[k-i] = a[k-i], a[k]
				}
			}
		}
	}

	return a
}
