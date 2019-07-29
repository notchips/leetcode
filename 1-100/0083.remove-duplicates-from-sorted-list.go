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
	headNode := &ListNode{0, nil}
	retNode := headNode
	var pre *ListNode
	for node := head; node != nil; {
		if node == head || node.Val != pre.Val {
			retNode.Next = node
			retNode = node
		}
		pre = node
		node = node.Next
	}
	retNode.Next = nil
	return headNode.Next
}
