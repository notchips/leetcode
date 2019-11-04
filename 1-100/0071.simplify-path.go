/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */
package leetcode

import (
	"strings"
)

// @lc code=start
func simplifyPath(path string) string {
	nodes := strings.Split(path, "/")
	var stack []string
	for _, node := range nodes {
		if node == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if len(node) > 0 && node != "." && node != ".." {
			stack = append(stack, node)
		}
	}
	return "/" + strings.Join(stack, "/")
}

// @lc code=end
