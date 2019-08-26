/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
func plusOne(digits []int) []int {
	n, carry := len(digits), 0
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			digits[i], carry = (digits[i]+1)%10, (digits[i]+1)/10
		} else {
			digits[i], carry = (digits[i]+carry)%10, (digits[i]+carry)/10
		}
	}
	if carry > 0 {
		return append([]int{carry}, digits...)
	}
	return digits
}
