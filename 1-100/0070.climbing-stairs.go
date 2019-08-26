/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
func climbStairs(n int) int {
	pre, post := 1, 2
	for i := 0; i < n-1; i++ {
		pre, post = post, pre+post
	}
	return pre
}
