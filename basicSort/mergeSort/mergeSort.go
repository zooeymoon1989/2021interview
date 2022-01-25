package mergeSort

func mergeSort(a []int) []int {

	if len(a) < 2 {
		return a
	}
	return conquer(mergeSort(a[0:len(a)/2]), mergeSort(a[len(a)/2:]))
}

func conquer(left, right []int) []int {
	var result []int
	var i, j = 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
			if j == len(right) {
				result = append(result, left[i:]...)
			}
		} else {
			result = append(result, left[i])
			i++
			if i == len(left) {
				result = append(result, right[j:]...)
			}
		}
	}
	return result
}
