/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
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

// 递归
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	if l1.Val <= l2.Val {
//		l1.Next = mergeTwoLists(l1.Next, l2)
//		return l1
//	}
//	l2.Next = mergeTwoLists(l1, l2.Next)
//	return l2
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	headNode := new(ListNode)
	pre := headNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1, pre = l1.Next, pre.Next
		} else {
			pre.Next = l2
			l2, pre = l2.Next, pre.Next
		}
	}
	if l1 != nil {
		pre.Next = l1
	} else {
		pre.Next = l2
	}
	return headNode.Next
}

// @lc code=end
