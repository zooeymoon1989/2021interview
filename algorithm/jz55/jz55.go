package jz55

import (
	. "interviewPractice/nc_tools"
	"math"
)

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pRoot TreeNode类
 * @return int整型
 */
func TreeDepth(pRoot *TreeNode) int {
	// write code here
	if pRoot == nil {
		return 0
	}
	l := TreeDepth(pRoot.Left)
	r := TreeDepth(pRoot.Right)
	return int(math.Max(float64(l), float64(r))) + 1
}
