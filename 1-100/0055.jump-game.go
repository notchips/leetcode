/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */
func canJump(nums []int) bool {
	for start, end := 0, len(nums)-1; start != end && start+nums[start] < end; {
		if nums[start] == 0 {
			return false
		}
		bestNext := start + 1
		for next := start + 2; next <= start+nums[start]; next++ {
			if next-bestNext+nums[next] >= nums[bestNext] {
				bestNext = next
			}
		}
		start = bestNext
	}
	return true
}
