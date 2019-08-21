/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {	return nil }
	if n == 1 { return lists[0] }
	return mergeTwoLists(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil { return l2 }
	if l2 == nil { return l1 }
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} 
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
