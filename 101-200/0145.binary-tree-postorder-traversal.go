/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var list []int
	var s []*TreeNode
	p := root
	for len(s) > 0 || p != nil {
		if p != nil {
			s = append(s, p)
			list = append(list, p.Val)
			p = p.Right
		} else {
			top := s[len(s)-1]
			s = s[:len(s)-1]
			p = top.Left
		}
	}
	reverse(list)
	return list
}

func reverse(a []int) {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
}
// @lc code=end

