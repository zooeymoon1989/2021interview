package jz21

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array int整型一维数组
 * @return int整型一维数组
 */
func reOrderArray(array []int) []int {
	// write code here
	if len(array) < 2 {
		return array
	}

	var i = 0
	for j := 0; j < len(array); j++ {
		if array[j]%2 == 1 {
			var tmp = array[j]
			for k := j - 1; k >= i; k-- {
				array[k], array[k+1] = array[k+1], array[k]
			}
			array[i] = tmp
			i++
		}
	}
	return array
}
