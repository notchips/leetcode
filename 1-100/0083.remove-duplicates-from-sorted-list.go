/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	headNode := &ListNode{}
	tempNode := headNode
	var pre *ListNode
	for head != nil {
		if !equal(pre, head) {
			tempNode.Next = head
			tempNode = tempNode.Next
		}
		pre, head = head, head.Next
	}
	tempNode.Next = nil
	return headNode.Next
}

func equal(a, b *ListNode) bool {
	return a != nil && b != nil && a.Val == b.Val
}
