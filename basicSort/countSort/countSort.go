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
	m := make([]int, max-min+1)
	var result []int
	for i := 0; i < len(a); i++ {
		m[a[i]-min]++
	}
	for i := 0; i < len(m); i++ {
		k := m[i]
		for k > 0 {
			result = append(result, i+min)
			k--
		}
	}
	return result
}
