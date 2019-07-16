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
	tail := 0
	var head, pre *ListNode = l1, nil
	for ; l1 != nil && l2 != nil; pre, l1, l2 = l1, l1.Next, l2.Next{
		sum := l1.Val + l2.Val + tail
		l1.Val, tail = sum % 10, sum / 10
	}
	for ; tail > 0 && l1 != nil; pre, l1 = l1, l1.Next{
		sum := l1.Val + tail
		l1.Val, tail = sum % 10, sum / 10
	}
	if l2 != nil{
		pre.Next = l2
		for ; tail > 0 && l2 != nil; pre, l2 = l2, l2.Next{
			sum := l2.Val + tail
			l2.Val, tail = sum % 10, sum / 10
		}
	}
	if tail != 0{
		pre.Next = &ListNode{
			Val: tail,
			Next: nil,
		}
	}
    return head
}
