/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 && len(newInterval) == 0 {
		return [][]int{}
	}
	ans := make([][]int, 0, len(intervals))
	for _, interval := range intervals {
		// newInterval在interval的前面或者后面
		if newInterval[1] < interval[0] || interval[1] < newInterval[0] {
			ans = append(ans, interval)
		} else { // newInterval和interval相交
			if interval[0] < newInterval[0] {
				newInterval[0] = interval[0]
			}
			if interval[1] > newInterval[1] {
				newInterval[1] = interval[1]
			}
		}
	}
	ans = append(ans, newInterval)
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}
