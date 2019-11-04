/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
package leetcode

// @lc code=start
func plusOne(digits []int) []int {
	n, carry := len(digits), 0
	for i := n - 1; i >= 0; i-- {
		var sum int
		if i == n-1 {
			sum = digits[i] + 1
		} else {
			sum = digits[i] + carry
		}
		digits[i], carry = sum%10, sum/10
	}
	if carry == 0 {
		return digits
	}
	return append([]int{1}, digits...)
}

// @lc code=end
