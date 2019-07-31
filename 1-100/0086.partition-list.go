/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	// 将原链表分为两部分，左部分小于x，右部分大或等于x
	leftHeadNode := &ListNode{0, nil}
	leftTempNode := leftHeadNode
	rightHeadNode := &ListNode{0, nil}
	rightTempNode := rightHeadNode
	for head != nil {
		if head.Val < x {
			leftTempNode.Next = head
			leftTempNode = head
		} else {
			rightTempNode.Next = head
			rightTempNode = head
		}
		head = head.Next
	}
	rightTempNode.Next = nil
	leftTempNode.Next = rightHeadNode.Next
	return leftHeadNode.Next
}
