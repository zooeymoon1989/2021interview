package heapSort

func heapSort(a []int) []int {
	heapSlice(a, 0)
	for i := len(a) - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapSlice(a[:i], 0)
	}
	return a
}

func heapSlice(a []int, offset int) {
	if offset*2+1 < len(a) && a[offset] <= a[offset*2+1] {
		a[offset], a[offset*2+1] = a[offset*2+1], a[offset]
	} else {
		return
	}

	if offset*2+2 < len(a) && a[offset] <= a[offset*2+2] {
		a[offset], a[offset*2+2] = a[offset*2+2], a[offset]
	} else {
		return
	}
	heapSlice(a, offset*2+1)
	heapSlice(a, offset*2+2)
}
