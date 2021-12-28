package heapSort

import "fmt"

func heapSort(a []int) []int {

	for i := 0; i < len(a)/2+1; i++ {
		heapSlice(a, i)
	}

	for i := len(a) - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapSlice(a[:i], 0)
	}
	fmt.Printf("%+v\n", a)
	return a
}

func heapSlice(a []int, offset int) {
	if offset*2+1 < len(a) && a[offset] <= a[offset*2+1] {
		a[offset], a[offset*2+1] = a[offset*2+1], a[offset]
		heapSlice(a, offset*2+1)
	}

	if offset*2+2 < len(a) && a[offset] <= a[offset*2+2] {
		a[offset], a[offset*2+2] = a[offset*2+2], a[offset]
		heapSlice(a, offset*2+2)
	}
}
