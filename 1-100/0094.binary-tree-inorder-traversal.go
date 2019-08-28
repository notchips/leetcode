/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// recursive 
// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}
// 	var values []int
// 	values = append(values, inorderTraversal(root.Left)...)
// 	values = append(values, root.Val)
// 	values = append(values, inorderTraversal(root.Right)...)
// 	return values
// }

// iterative
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{root}
	values := make([]int, 0, 16)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top.Left != nil {
			stack = append(stack, top.Left)
			top.Left = nil
		} else {
			values = append(values, top.Val)
			stack = stack[:len(stack)-1]
			if top.Right != nil {
				stack = append(stack, top.Right)
				top.Right = nil
			}
		}
	}
	return values
}
