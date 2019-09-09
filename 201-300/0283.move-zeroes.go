/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */
func moveZeroes(nums []int) {
	pre := 0
	// 首先让pre指向第一个0
	for pre < len(nums) && nums[pre] != 0 {
		pre++
	}
	post := pre + 1
	// 首先让post指向Pre后的第一个非0
	for post < len(nums) && nums[post] == 0 {
		post++
	}

	for post < len(nums) {
		nums[pre], nums[post] = nums[post], nums[pre]
		for pre < len(nums) && nums[pre] != 0 {
			pre++
		}
		for post < len(nums) && nums[post] == 0 {
			post++
		}
	}
}
