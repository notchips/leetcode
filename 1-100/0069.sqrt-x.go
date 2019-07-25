/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
func mySqrt(x int) int {
	left, right := 0, x
	for left < right {
		mid := left + (right-left)/2 + 1
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
