package jz15

func NumberOf1(i int) int {
	// write code here
	count := 0

	for i != 0 {
		count++
		i = i & int(int32(i)-1)
	}
	return count
}
