/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n < 1 {
		return nil
	}
	slow, fast := head, head
	for fast != nil && n > 0 {
		fast = fast.Next
		n--
	}
	if n > 0 {
		return nil
	}
	headNode := &ListNode{Next: slow}
	pre := headNode
	for fast != nil {
		slow, fast, pre = slow.Next, fast.Next, slow
	}
	pre.Next = slow.Next
	return headNode.Next
}

// @lc code=end
