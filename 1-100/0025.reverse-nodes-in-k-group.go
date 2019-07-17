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
	headNode := &ListNode{0, head}
	pre := headNode
	nodes := make([]*ListNode, 0, k)
	for head != nil {
		nodes = append(nodes, head)
		if len(nodes) == k {
			pre.Next = nodes[k-1]
			pre = nodes[0]
			nodes[0].Next = nodes[k-1].Next
			for i := k - 1; i > 0; i-- {
				nodes[i].Next = nodes[i-1]
			}
			nodes = nodes[:0]
			head = pre.Next
		} else {
			head = head.Next
		}
		
	}
	return headNode.Next
}
