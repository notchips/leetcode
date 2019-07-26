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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := l1
	var pre *ListNode // l1 pre-node
	for l1 != nil && l2 != nil {
		l1.Val, carry = (l1.Val+l2.Val+carry)%10, (l1.Val+l2.Val+carry)/10
		pre, l1, l2 = l1, l1.Next, l2.Next
	}
	if l2 != nil { // l1 == nil && l2 !=nil
		pre.Next = l2
		for carry > 0 && l2 != nil {
			l2.Val, carry = (l2.Val+carry)%10, (l2.Val+carry)/10
			pre, l2 = l2, l2.Next
		}
	} else { // l1 != nil && l2 == nil
		for carry > 0 && l1 != nil {
			l1.Val, carry = (l1.Val+carry)%10, (l1.Val+carry)/10
			pre, l1 = l1, l1.Next
		}
	}
	if carry != 0 {
		pre.Next = &ListNode{carry, nil}
	}
	return head
}
