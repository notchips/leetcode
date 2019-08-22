/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		f := myPow(x, n/2)
		return f * f
	}
	if n < 0 {
		return 1 / x * myPow(x, n+1)
	}
	return x * myPow(x, n-1)
}
