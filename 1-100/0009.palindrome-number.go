/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == reverse(x)
}

func reverse(x int) int {
	y := 0
	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}
	return y
}
