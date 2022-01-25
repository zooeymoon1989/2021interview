package nnnnn

func Nnnnnn(a []int, b int) int {
	if len(a) == 0 {
		return -1
	}

	min := 0
	max := len(a) - 1
	for min <= max {
		mid := min + (max-min)/2
		if a[mid] > b {
			max = mid - 1
		} else if a[mid] < b {
			min = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
