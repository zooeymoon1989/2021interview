package jz27

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pRoot TreeNode类
 * @return TreeNode类
 */
func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here
	if pRoot == nil {
		return pRoot
	}
	pRoot.Left, pRoot.Right = Mirror(pRoot.Right), Mirror(pRoot.Left)
	return pRoot
}
