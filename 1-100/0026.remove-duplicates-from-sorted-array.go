/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
package leetcode

// @lc code=start
func removeDuplicates(nums []int) int {
	p := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[p] = nums[i]
			p++
		}
	}
	return p
}

// @lc code=end
