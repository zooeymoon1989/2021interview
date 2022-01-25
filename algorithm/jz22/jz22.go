package jz22

import (
	. "interviewPractice/nc_tools"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pHead ListNode类
 * @param k int整型
 * @return ListNode类
 */
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}
	var fast, slow = pHead, pHead
	for i := 0; i < k; i++ {
		if fast != nil {
			fast = fast.Next
		} else {
			return nil
		}
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow

}
