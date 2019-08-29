/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	balanced := true
	checkHeight(root, &balanced)
	return balanced
}

func checkHeight(root *TreeNode, balanced *bool) int {
	if root == nil {
		return 0
	}
	leftHeight := checkHeight(root.Left, balanced)
	rightHeight := checkHeight(root.Right, balanced)
	if abs(leftHeight-rightHeight) > 1 {
		*balanced = false
	}
	return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
