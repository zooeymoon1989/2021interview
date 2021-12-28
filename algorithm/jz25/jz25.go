package jz25

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	var result = new(ListNode)
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val > pHead2.Val {
			if result == nil {
				result.Val = pHead2.Val
			} else {
				result.Next = pHead2
			}
		} else {
			if result == nil {
				result.Val = pHead1.Val
			} else {
				result.Next = pHead1
			}
		}
	}

	return result
}
