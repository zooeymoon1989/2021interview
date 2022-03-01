package jz32

import . "interviewPractice/nc_tools"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param root TreeNode类
 * @return int整型一维数组
 */
func PrintFromTopToBottom(root *TreeNode) []int {
	// write code here
	var result []int
	if root == nil {
		return result
	}
	parent := []*TreeNode{root}
	for parent != nil {
		var child []*TreeNode
		for i := 0; i < len(parent); i++ {
			result = append(result, parent[i].Val)
			if parent[i].Left != nil {
				child = append(child, parent[i].Left)
			}
			if parent[i].Right != nil {
				child = append(child, parent[i].Right)
			}
		}
		parent = child
	}
	return result
}
