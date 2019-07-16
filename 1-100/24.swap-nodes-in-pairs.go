/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
	headNode := &ListNode{0, head}
	pre := headNode
	for head != nil {
		if head.Next != nil {
			pre.Next = head.Next
			pre = head
			head.Next, head.Next.Next = head.Next.Next, head
		}
		head = head.Next
	}
	return headNode.Next
}
