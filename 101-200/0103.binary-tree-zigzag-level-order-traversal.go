/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	return bfs([]*TreeNode{root}, 0)
}

func bfs(roots []*TreeNode, level int) [][]int {
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
	if level&1 == 1 {
		reverse(values)
	}
	return append([][]int{values}, bfs(children, level+1)...)
}

func reverse(values []int) {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
}
