/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
package leetcode

// @lc code=start
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	pre, post := 1, 1
	for ; n > 0; n-- {
		pre, post = post, pre+post
	}
	return pre
}

// @lc code=end
