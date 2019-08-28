/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return []*TreeNode{}
	}
	return buildTrees(1, n)
}

func buildTrees(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}
	var rootList []*TreeNode
	for mid := left; mid <= right; mid++ {
		leftChildTrees := buildTrees(left, mid-1)
		rightChildTrees := buildTrees(mid+1, right)
		for _, leftChildTree := range leftChildTrees {
			for _, rightChildTree := range rightChildTrees {
				root := &TreeNode{
					Val:   mid,
					Left:  leftChildTree,
					Right: rightChildTree,
				}
				rootList = append(rootList, root)
			}
		}
	}
	return rootList
}
