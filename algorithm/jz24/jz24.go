package jz24

import . "interviewPractice/nc_tools"

/**
 *
 * @param pHead ListNodeš▒╗
 * @return ListNodeš▒╗
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here

	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	var pre, cur, next *ListNode
	cur = pHead
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
