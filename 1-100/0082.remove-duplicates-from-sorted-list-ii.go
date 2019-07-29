/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	headNode := &ListNode{0, nil} 
	retPre := headNode  
	var pre *ListNode  // 记录node的前置节点
	for node := head; node != nil; {
		// 如果node满足： 
		// 1.同时是头节点和尾节点 或者
		// 2.是头节点且不等于后置节点的值 或者
		// 3.不等于前置节点的值且是尾节点 或者 
		// 4.不等于前后节点的值
		// 则添加到结果中
		if (node == head || node.Val != pre.Val) && (node.Next == nil || node.Val != node.Next.Val) {
			retPre.Next = node
			retPre = node
		}
		pre = node
		node = node.Next
	}
	retPre.Next = nil
	return headNode.Next
}
