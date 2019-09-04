/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */
func majorityElement(nums []int) int {
	num, cnt := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			num, cnt = nums[i], 1
		} else if num == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return num
}
