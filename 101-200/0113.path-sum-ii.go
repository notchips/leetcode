/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var paths [][]int
	path := []int{root.Val}
	dfs(root, sum-root.Val, &paths, &path)
	return paths
}

func dfs(root *TreeNode, sum int, paths *[][]int, path *[]int) {
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			newPath := make([]int, len(*path))
			copy(newPath, *path)
			*paths = append(*paths, newPath)
		}
		return
	}
	for _, child := range []*TreeNode{root.Left, root.Right} {
		if child != nil {
			*path = append(*path, child.Val)
			dfs(child, sum-child.Val, paths, path)
			*path = (*path)[:len(*path)-1]
		}
	}
}
