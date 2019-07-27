/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */
func sortColors(nums []int) {
	var cnt [2]int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 || nums[i] == 1 {
			cnt[nums[i]]++
		}
	}
	for i := 0; i < len(nums); i++ {
		if cnt[0] > 0 {
			nums[i] = 0
			cnt[0]--
		} else if cnt[1] > 0 {
			nums[i] = 1
			cnt[1]--
		} else {
			nums[i] = 2
		}
	}
}
