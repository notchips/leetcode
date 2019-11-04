/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
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
func rotateRight(head *ListNode, k int) *ListNode {
	// 计算链表长度，同时找到尾结点
	var tail *ListNode
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		tail = cur
		length++
	}

	if length == 0 || k%length == 0 {
		return head
	}

	// 寻找新的头结点和尾结点
	var newHead, newTail *ListNode
	k = length - k%length
	for cur := head; cur != nil; cur = cur.Next {
		if k == 1 {
			newHead, newTail = cur.Next, cur
			break
		}
		k--
	}
 
	tail.Next = head
	newTail.Next = nil
	return newHead
}

// @lc code=end
