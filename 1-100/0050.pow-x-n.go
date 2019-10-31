/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
package leetcode

// @lc code=start
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		t := myPow(x, n/2)
		return t * t
	}
	if n < 0 {
		return 1 / x * myPow(x, n+1)
	}
	return x * myPow(x, n-1)
}

// @lc code=end
