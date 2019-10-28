/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
package leetcode

// @lc code=start
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

// @lc code=end
