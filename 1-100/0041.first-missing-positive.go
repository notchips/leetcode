/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */
package leetcode

// @lc code=start
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for 0 <= nums[i]-1 && nums[i]-1 < len(nums) &&
			nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			return i + 1
		}
	}
	return len(nums) + 1
}

// @lc code=end
