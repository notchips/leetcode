/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
package leetcode

import "math"

// @lc code=start
func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x /= 10
		if y < math.MinInt32 || y > math.MaxInt32 {
			return 0
		}
	}
	return y
}

// @lc code=end
