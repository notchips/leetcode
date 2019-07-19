/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */
 func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		min := math.MaxInt32
		for j := i + 1; j < n && j <= i+nums[i]; j++ {
			if dp[j] < min {
				min = dp[j]
			}
		}
		dp[i] = 1 + min
	}
	return dp[0]
}
