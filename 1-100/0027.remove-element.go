/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */
package leetcode

// @lc code=start
func removeElement(nums []int, val int) int {
	p := 0
	for _, num := range nums {
		if num != val {
			nums[p] = num
			p++
		}
	}
	return p
}

// @lc code=end
