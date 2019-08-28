/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	values := make([]int, 0, 32)
	inOrder(root, &values)
	for i := 1; i < len(values); i++ {
		if values[i] <= values[i-1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, values)
	*values = append(*values, root.Val)
	inOrder(root.Right, values)
}
