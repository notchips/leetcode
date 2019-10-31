/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
package leetcode

import "math"

// @lc code=start
func maxSubArray(nums []int) int {
	maxSum, sum := math.MinInt64, 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

// @lc code=end
