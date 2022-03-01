package jz7

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
 * @param pre int整型一维数组
 * @param vin int整型一维数组
 * @return TreeNode类
 */
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here
	if len(pre) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   pre[0],
		Left:  nil,
		Right: nil,
	}

	var mid int

	for i := 0; i < len(pre); i++ {
		if vin[i] == pre[0] {
			mid = i
			break
		}
	}

	root.Left = reConstructBinaryTree(pre[1:mid+1], vin[:mid])
	root.Right = reConstructBinaryTree(pre[mid+1:], vin[mid+1:])
	return root
}
