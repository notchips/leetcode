/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */
package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//递归
//func swapPairs(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	next := head.Next
//	head.Next = swapPairs(next.Next)
//	next.Next = head
//	return next
//}

func swapPairs(head *ListNode) *ListNode {
	headNode := new(ListNode)
	pre := headNode
	// pre -> head -> next -> post
	for head != nil {
		next := head.Next
		if next == nil {
			pre.Next = head
			head, pre = head.Next, pre.Next
		} else {
			post := next.Next
			pre.Next, next.Next = next, head
			head, pre = post, head
		}
	}
	pre.Next = nil
	return headNode.Next
}

// @lc code=end
