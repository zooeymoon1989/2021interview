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
	i := 1
	j := len(a) - 1
	for i < j {

	}
	return a
}

func pivot(n int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(n)
}
