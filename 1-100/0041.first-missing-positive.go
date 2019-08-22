/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */
func firstMissingPositive(nums []int) int {
	nums = append(nums, -1)
	n := len(nums)
	for i := 0; i < n; i++ {
		for 0 <= nums[i] && nums[i] < n && nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for i := 1; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	return n
}
