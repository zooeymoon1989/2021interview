package jz77

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
 * @param pRoot TreeNode类
 * @return int整型二维数组
 */
func Print(pRoot *TreeNode) [][]int {
	// write code here
	var result [][]int
	if pRoot == nil {
		return result
	}
	var level = 1
	parent := []*TreeNode{pRoot}
	for len(parent) > 0 {
		var r []int
		var child []*TreeNode
		for i := 0; i < len(parent); i++ {
			if parent[i].Left != nil {
				child = append(child, parent[i].Left)
			}
			if parent[i].Right != nil {
				child = append(child, parent[i].Right)
			}
		}

		if level%2 == 1 {
			for i := 0; i < len(parent); i++ {
				r = append(r, parent[i].Val)
			}
		} else {
			for i := len(parent) - 1; i >= 0; i-- {
				r = append(r, parent[i].Val)
			}
		}

		result = append(result, r)
		parent = child
		level++
	}
	return result
}
