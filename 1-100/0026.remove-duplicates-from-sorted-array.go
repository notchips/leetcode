/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
 func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[cnt] = nums[i]
			cnt++
		}
	}
	return cnt
}
