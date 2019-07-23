/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */
func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	var ans [][]int
	stage := 0
	var left int
	for i := 0; i < n; i++ {
		if stage == 2 {
			ans = append(ans, intervals[i])
		}

		if stage == 0 {
			if intervals[i][1] < newInterval[0] {
				ans = append(ans, intervals[i])
			} else {
				left = min(intervals[i][0], newInterval[0])
				stage = 1
			}
		}

		if stage == 1 {
			if newInterval[1] <= intervals[i][1] {
				if newInterval[1] >= intervals[i][0] {
					ans = append(ans, []int{left, intervals[i][1]})
				} else {
					ans = append(ans, []int{left, newInterval[1]}, intervals[i])
				}
				stage = 2
			} else if i == n-1 {
				ans = append(ans, []int{left, newInterval[1]})
			}
		}
	}

	if stage == 0 {
		ans = append(ans, newInterval)
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
