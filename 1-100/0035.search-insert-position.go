/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
func searchInsert(nums []int, target int) int {
	lo, hi, pos := 0, len(nums)-1, -1
	for lo <= hi {
		md := lo + (hi-lo)/2
		if target == nums[md] {
			return md
		} else if target < nums[md] {
			hi = md - 1
		} else {
			lo = md + 1
			pos = max(pos, md)
		}
	}
	return pos + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
