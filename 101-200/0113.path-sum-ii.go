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
		return [][]int{}
	}
	paths := make([][]int, 0, 8)
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
	if root.Left != nil {
		*path = append(*path, root.Left.Val)
		dfs(root.Left, sum-root.Left.Val, paths, path)
		*path = (*path)[:len(*path)-1]
	}
	if root.Right != nil {
		*path = append(*path, root.Right.Val)
		dfs(root.Right, sum-root.Right.Val, paths, path)
		*path = (*path)[:len(*path)-1]
	}
}
