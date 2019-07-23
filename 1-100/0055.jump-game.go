/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */
func canJump(nums []int) bool {
	n := len(nums)
	for i := 0; i < n-1; {
		if nums[i] == 0 {
			return false
		}

		// 贪心，选择能跳最大距离的下一跳位置，同等距离情况下选择靠右的
		max := i + 1
		for j := i + 2; j < n && j <= i+nums[i]; j++ {
			if nums[j]+j >= nums[max]+max {
				max = j
			}
		}
		i = max
	}
	return true
}
