/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
	nodes := make([]*ListNode, 0, k)
	for temp, i := head, 0; temp != nil && i < k; i++ {
		nodes = append(nodes, temp)
		temp = temp.Next
	}
	if len(nodes) < k { return head }
	nodes[0].Next = reverseKGroup(nodes[k-1].Next, k)
	for i := k - 1; i > 0; i-- {
		nodes[i].Next = nodes[i-1]
	}
	return nodes[k-1]
}
