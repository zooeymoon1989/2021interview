package quickSort

import (
	"math/rand"
	"time"
)

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	p := pivot(len(a))
	a[0], a[p] = a[p], a[0]
	var left []int
	var right []int
	for i := 1; i < len(a); i++ {
		if a[i] > a[0] {
			right = append(right, a[i])
		} else {
			left = append(left, a[i])
		}
	}
	return append(quickSort(append(left, a[0])), quickSort(right)...)
}

func pivot(n int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(n)
}
