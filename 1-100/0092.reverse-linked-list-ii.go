/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	headNode := &ListNode{Next: head} // 带头结点链表
	preNode,curNode := headNode, head // 指向每次遍历时的前置结点和当前结点
	var beforeReverseNode, firstReverseNode *ListNode
	for i := 1; i <= n; i++ {
		if i == m { // 记录翻转之前的结点和第一个翻转结点
			beforeReverseNode, firstReverseNode = preNode, curNode
		}
		postNode := curNode.Next // 记录后置结点
		if i > m && i <= n { // 在(m, n]范围内，当前结点修改为指向自己的前置结点
			curNode.Next = preNode
		}
		if i == n { // 当遍历到翻转的最后一个结点时，进行收尾工作
			beforeReverseNode.Next = curNode
			firstReverseNode.Next = postNode
		}
		// 更新前置结点和当前结点
		preNode = curNode
		curNode = postNode
	}
	return headNode.Next
}
