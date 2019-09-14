/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
package leetcode

// @lc code=start
func isPalindrome(x int) bool {
	y := 0
	for t := x; t > 0; t /= 10 {
		y = y*10 + t%10
	}
	return x == y
}

// @lc code=end
