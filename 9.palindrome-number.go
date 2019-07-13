/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)
	for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1{
		if s[i] != s[j] {
			return false
		}
	 }
	 return true
}

