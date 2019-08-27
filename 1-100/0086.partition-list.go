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
	leftHeadNode, rightHeadNode := new(ListNode), new(ListNode)
	leftPreNode, rightPreNode := leftHeadNode, rightHeadNode
	for ; head != nil; head = head.Next {
		if head.Val < x {
			leftPreNode.Next = head
			leftPreNode = leftPreNode.Next
		} else {
			rightPreNode.Next = head
			rightPreNode = rightPreNode.Next
		}
	}
	rightPreNode.Next = nil
	leftPreNode.Next = rightHeadNode.Next
	return leftHeadNode.Next
}
