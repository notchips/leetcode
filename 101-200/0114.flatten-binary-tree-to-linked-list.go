/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	nodes := make([]*TreeNode, 0, 32)
	preOrder(root, &nodes)
	n := len(nodes)
	if n == 0 {
		return
	}
	for i := 0; i < n-1; i++ {
		nodes[i].Left, nodes[i].Right = nil, nodes[i+1]
	}
	nodes[n-1].Left, nodes[n-1].Right = nil, nil
}

func preOrder(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	*nodes = append(*nodes, root)
	preOrder(root.Left, nodes)
	preOrder(root.Right, nodes)
}
