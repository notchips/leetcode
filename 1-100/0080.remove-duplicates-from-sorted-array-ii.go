/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */
func removeDuplicates(nums []int) int {
	pos := 0
	for _, num := range nums {
		if pos < 2 || num != nums[pos-2] {
			nums[pos] = num
			pos++
		}
	}
	return pos
}
