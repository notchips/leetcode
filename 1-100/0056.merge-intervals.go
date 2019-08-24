/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0, 16)
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > right {
			ans = append(ans, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else if intervals[i][1] > right {
			right = intervals[i][1]
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}
