/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */
func jump(nums []int) int {
	cnt := 0
	for start, end := 0, len(nums)-1; start != end; cnt++ {
		if start+nums[start] >= end {
			cnt++
			break
		}
		bestNext := start + 1
		for next := start + 2; next <= start+nums[start]; next++ {
			if next-bestNext+nums[next] >= nums[bestNext] {
				bestNext = next
			}
		}
		start = bestNext
	}
	return cnt
}
