/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */
package leetcode

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := make([][]int, 0, len(intervals))
	inserted := false
	for _, interval := range intervals {
		if !inserted {
			if newInterval[1] < interval[0] {
				ans = append(ans, newInterval, interval)
				inserted = true
			} else if interval[1] < newInterval[0] {
				ans = append(ans, interval)
			} else {
				newInterval[0] = min(newInterval[0], interval[0])
				newInterval[1] = max(newInterval[1], interval[1])
			}
		} else {
			ans = append(ans, interval)
		}
	}
	if !inserted {
		ans = append(ans, newInterval)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
