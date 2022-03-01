package jz26

import . "interviewPractice/nc_tools"

func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	// write code here
	if pRoot2 == nil || pRoot1 == nil {
		return false
	}
	return isSubtree(pRoot1, pRoot2) || HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
}

func isSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot2 == nil {
		return true
	}

	if pRoot1 == nil {
		return false
	}
	if pRoot1.Val != pRoot2.Val {
		return false
	}
	return isSubtree(pRoot1.Left, pRoot2.Left) && isSubtree(pRoot1.Right, pRoot2.Right)
}
