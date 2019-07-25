/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	length := 0
	var tail *ListNode
	for temp := head; temp != nil; temp = temp.Next {
		length++
		tail = temp
	}
	k %= length

	if length == 1 || k == 0 {
		return head
	}

	tail.Next = head

	newTail := head
	for i := 1; i < length-k; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}
