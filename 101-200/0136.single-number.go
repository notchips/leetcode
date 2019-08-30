/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */
func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
