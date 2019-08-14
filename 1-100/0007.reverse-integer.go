/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
 func reverse(x int) int {
	ans := 0
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}
	return ans
}
