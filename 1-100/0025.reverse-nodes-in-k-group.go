/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	headNode := new(ListNode) // 带头结点链表
	preNode := headNode       // 前k个结点的最后一个结点，即当前k个结点的前置结点
	for {
		first := head
		cnt := k
		for head != nil && cnt > 0 {
			head = head.Next
			cnt--
		}
		if cnt > 0 { // 当前区间不满k个，跳出循环
			preNode.Next = first
			break
		}
		// 将[first, head)区间的结点，按头插法插入到preNode后
		for cur := first; cur != head; {
			preNode.Next, cur.Next, cur = cur, preNode.Next, cur.Next
		}
		preNode = first
	}
	return headNode.Next
}

// @lc code=end
