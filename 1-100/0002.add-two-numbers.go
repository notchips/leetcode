/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	headNode := &ListNode{Next: l1}
	pre := headNode
	carry := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		l1.Val, carry = sum%10, sum/10
		l1, l2, pre = l1.Next, l2.Next, l1
	}
	if l2 != nil { // l1 == nil && l2 != nil
		l1, pre.Next = l2, l2
	}
	for carry > 0 && l1 != nil { // l1 != nil && l2 == nil
		sum := l1.Val + carry
		l1.Val, carry = sum%10, sum/10
		l1, pre = l1.Next, l1
	}
	if carry > 0 {
		pre.Next = &ListNode{Val: carry, Next: nil}
	}
	return headNode.Next
}

// @lc code=end
