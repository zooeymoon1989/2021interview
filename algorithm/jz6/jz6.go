package jz6

import (
	. "interviewPractice/nc_tools"
)

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	var result []int
	cur := head
	for cur != nil {
		result = append(result, cur.Val)
		//result = append([]int{(*cur).Val}, result...)
		cur = cur.Next
	}
	return result
}
