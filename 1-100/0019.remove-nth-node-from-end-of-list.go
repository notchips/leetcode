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
	headNode := &ListNode{Next: head}
	nodes := make([]*ListNode, 0, 10)
	tempNode := headNode
	for tempNode != nil {
		nodes = append(nodes, tempNode)
		tempNode = tempNode.Next
	}
	target := len(nodes) - n
	nodes[target-1].Next = nodes[target].Next
	return headNode.Next
}
