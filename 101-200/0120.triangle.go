/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */
func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	for row := 0; row < len(triangle); row++ {
		pre := 0 // 记录上一层同一位置左边的值
		for col := 0; col < len(triangle[row]); col++ {
			if col == 0 {
				dp[col], pre = triangle[row][col]+dp[col], dp[col]
			} else if 0 < col && col < len(triangle[row])-1 {
				dp[col], pre = triangle[row][col]+min(pre, dp[col]), dp[col]
			} else {
				dp[col] = triangle[row][col] + pre
			}
		}
	}
	return min(dp...)
}

func min(a ...int) int {
	ret := a[0]
	for _, v := range a {
		if v < ret {
			ret = v
		}
	}
	return ret
}
