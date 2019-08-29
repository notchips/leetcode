/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return bfs([]*TreeNode{root})
}

func bfs(roots []*TreeNode) int {
	children := make([]*TreeNode, 0, 2*len(roots))
	for _, root := range roots {
		if root.Left == nil && root.Right == nil {
			return 1
		}
		if root.Left != nil {
			children = append(children, root.Left)
		}
		if root.Right != nil {
			children = append(children, root.Right)
		}
	}
	return bfs(children) + 1
}
