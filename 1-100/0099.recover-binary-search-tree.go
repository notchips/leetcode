/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	nodes := make([]*TreeNode, 0, 32)
	inOrder(root, &nodes)
	for i := 0; i < len(nodes)-1; i++ {
		if nodes[i].Val > nodes[i+1].Val {
			for j := len(nodes) - 1; j > i; j-- {
				if nodes[j].Val < nodes[j-1].Val {
					nodes[i].Val, nodes[j].Val = nodes[j].Val, nodes[i].Val
					return
				}
			}
		}
	}
}

func inOrder(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left, nodes)
	*nodes = append(*nodes, root)
	inOrder(root.Right, nodes)
}
