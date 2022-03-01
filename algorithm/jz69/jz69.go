package jz69

func jumpFloor(number int) int {
	// write code here
	a, b, c := 1, 1, 1
	for i := 2; i <= number; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
