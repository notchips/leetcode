/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
func longestPalindrome(s string) string {
	 if len(s) < 2 {
		 return s
	 }
	 var start, maxLen int
	 for i := 0; i < len(s)-1; i++ {
		 s1, l1 := extend(i, i, s)  // 奇数长度扩展
		 if l1 > maxLen {
			 start, maxLen = s1, l1
		 }
		 s2, l2 := extend(i, i+1, s) // 偶数长度扩展
		 if l2 > maxLen {
			 start, maxLen = s2, l2
		 }
	 }
	 return s[start : start+maxLen]
 }
 
 func extend(i, j int, s string) (start int, len int) {
	 for i >= 0 && j < len(s) && s[i] == s[j] {
		 i, j = i-1, j+1
	 }
	 start, len = i+1, j-i-1
	 return
 }
