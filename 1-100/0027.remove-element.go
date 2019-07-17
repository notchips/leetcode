/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */
func removeElement(nums []int, val int) int {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}
