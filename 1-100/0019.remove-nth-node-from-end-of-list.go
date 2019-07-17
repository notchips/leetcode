/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptrs := make([]*ListNode, 0, 10)
	for ; head != nil; head = head.Next {
		ptrs = append(ptrs, head)
	}
	if len(ptrs) == n { // delete head node
		ptrs = append(ptrs, nil)
		return ptrs[1]
	}
	ptrs[len(ptrs)-n-1].Next = ptrs[len(ptrs)-n].Next
	return ptrs[0]
}
