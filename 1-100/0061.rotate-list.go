/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	nodes := make([]*ListNode, 0,  16)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	n := len(nodes)
	k %= n
	if k == 0 {
		return nodes[0]
	}
	nodes[n-1].Next = nodes[0]
	nodes[n-k-1].Next = nil
	return nodes[n-k]
}
