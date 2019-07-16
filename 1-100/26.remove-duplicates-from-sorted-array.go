/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
 func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pos := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[pos-1] {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}
