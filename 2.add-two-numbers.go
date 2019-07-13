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
	var (
		tail = 0
		head *ListNode = l1 // 利用l1拼接出结果
		pre *ListNode = nil // 记录l1的前置节点
	)
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

