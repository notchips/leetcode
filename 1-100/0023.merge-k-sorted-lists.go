/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
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
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	return mergeTwo(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func mergeTwo(l1, l2 *ListNode) *ListNode {
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
	if l1 == nil {
		pre.Next = l2
	} else {
		pre.Next = l1
	}
	return headNode.Next
}

// @lc code=end
