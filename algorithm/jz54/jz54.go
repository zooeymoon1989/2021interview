package jz54

import . "interviewPractice/nc_tools"

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
 *
 * @param proot TreeNode类
 * @param k int整型
 * @return int整型
 */
var result = -1
var tmp = 0

func KthNode(proot *TreeNode, k int) int {
	// write code here
	if proot == nil || k <= 0 {
		return -1
	}
	// write code here
	KthNode(proot.Left, k)
	tmp++
	if k == tmp {
		result = proot.Val
		return proot.Val
	}
	KthNode(proot.Right, k)
	return result
}
