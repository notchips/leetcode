/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */
package leetcode

import "sort"

// @lc code=start
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0, 16)
	left, right := intervals[0][0], intervals[0][1]
	for i := 0; i <= n; i++ {
		if i == n {
			ans = append(ans, []int{left, right})
		} else if right < intervals[i][0] {
			ans = append(ans, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else if intervals[i][1] > right {
			right = intervals[i][1]
		}
	}
	return ans
}

// @lc code=end
