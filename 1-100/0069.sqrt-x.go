/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
package leetcode

// @lc code=start
func mySqrt(x int) int {
	left, right := 0, x
	for left < right {
		mid := (left+right)/2 + 1
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}

// @lc code=end
