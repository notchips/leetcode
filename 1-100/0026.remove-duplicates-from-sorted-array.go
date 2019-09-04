/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
func removeDuplicates(nums []int) int {
	pos := 0
	for i, num := range nums {
		if i == 0 || num != nums[pos-1] {
			nums[pos] = num
			pos++
		}
	}
	return pos
}
