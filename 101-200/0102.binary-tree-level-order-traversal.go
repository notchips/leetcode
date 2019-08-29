/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	return bfs([]*TreeNode{root})
}

func bfs(roots []*TreeNode) [][]int {
	if len(roots) == 0 {
		return [][]int{}
	}
	values := make([]int, 0, len(roots))
	children := make([]*TreeNode, 0, 2*len(roots))
	for _, root := range roots {
		values = append(values, root.Val)
		if root.Left != nil {
			children = append(children, root.Left)
		}
		if root.Right != nil {
			children = append(children, root.Right)
		}
	}
	return append([][]int{values}, bfs(children)...)
}
