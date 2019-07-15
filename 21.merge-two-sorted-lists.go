/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head = &ListNode{0, nil}
	temp := head
	for l1 != nil && l2 != nil {
		if l2.Val < l1.Val {
			temp.Next = l2
			temp = temp.Next
			l2 = l2.Next
		} else {
			temp.Next = l1
			temp = temp.Next
			l1 = l1.Next
		}
	}
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return head.Next
}
