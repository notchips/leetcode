/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
func reverse(x int) int {
	var ret int
	for x != 0{
		ret = ret*10 + x%10
		x/=10
	}
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}
	return ret
}
