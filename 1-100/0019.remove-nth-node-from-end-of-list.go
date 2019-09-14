/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headNode := &ListNode{Next: head} // 带头结点
	pre := headNode // 指向移除结点的前结点
	for temp := head; temp != nil; temp, n = temp.Next, n-1 {
		if n <= 0 {
			pre = pre.Next
		}
	}
	pre.Next = pre.Next.Next
	return headNode.Next
}
