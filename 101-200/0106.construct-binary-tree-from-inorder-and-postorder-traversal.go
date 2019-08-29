/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	value := postorder[len(postorder)-1]
	var mid int
	for i, v := range inorder {
		if v == value {
			mid = i
			break
		}
	}
	return &TreeNode{
		Val:   value,
		Left:  buildTree(inorder[:mid], postorder[:mid]),
		Right: buildTree(inorder[mid+1:], postorder[mid:len(postorder)-1]),
	}
}
