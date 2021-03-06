package jz76

import . "interviewPractice/nc_tools"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func deleteDuplication(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	if pHead.Val != pHead.Next.Val {
		pHead.Next = deleteDuplication(pHead.Next)
		return pHead
	} else {
		tmp := pHead
		for tmp != nil && tmp.Val == pHead.Val {
			tmp = tmp.Next
		}
		return deleteDuplication(tmp)
	}
}
