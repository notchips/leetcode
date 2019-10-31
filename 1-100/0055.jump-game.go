/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */
package leetcode

// @lc code=start
func canJump(nums []int) bool {
	start, end := 0, len(nums)-1
	for start+nums[start] < end {
		if nums[start] == 0 {
			return false
		}
		bestNext := start + 1
		for next := start + 2; next <= start+nums[start]; next++ {
			if nums[next]+next > nums[bestNext]+bestNext {
				bestNext = next
			}
		}
		start = bestNext
	}
	return true
}

// @lc code=end
