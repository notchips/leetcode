/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 /*
  时间复杂度：O(n)；空间复杂度：O(1)
  √ 1563/1563 cases passed (8 ms)
  √ Your runtime beats 91.59 % of golang submissions
  √ Your memory usage beats 99.41 % of golang submissions (4.7 MB)
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head, pre *ListNode = l1, nil // pre记录l1的前置结点位置
	for l1 != nil && l2 != nil {
		l1.Val, carry = (l1.Val+l2.Val+carry)%10, (l1.Val+l2.Val+carry)/10
		pre, l1, l2 = l1, l1.Next, l2.Next
	}
	for carry > 0 && l1 != nil {
		l1.Val, carry = (l1.Val+carry)%10, (l1.Val+carry)/10
		pre, l1 = l1, l1.Next
	}
	if l2 != nil {
		pre.Next = l2
		for carry > 0 && l2 != nil {
			l2.Val, carry = (l2.Val+carry)%10, (l2.Val+carry)/10
			pre, l2 = l2, l2.Next
		}
	}
	if carry != 0 {
		pre.Next = &ListNode{carry, nil}
	}
	return head
}
