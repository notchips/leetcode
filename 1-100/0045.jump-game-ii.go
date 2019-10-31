/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */
package leetcode

// @lc code=start
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	cnt := 0
	for start, end := 0, len(nums)-1; start < end; cnt++ {
		if start+nums[start] >= end { // 能直接跳到最后一步
			start = end
		} else {
			bestNext := start + 1
			for next := start + 2; next <= start+nums[start]; next++ {
				if nums[next]+next >= nums[bestNext]+bestNext {
					bestNext = next
				}
			}
			start = bestNext
		}
	}
	return cnt
}

// @lc code=end
