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
	for i := 0; i < len(s); i++ { // 奇扩展
		s, l := extend(i, i, s) 
		if l > maxLen {
			start, maxLen = s, l
		}
	}
	for i := 0; i < len(s)-1; i++ { // 偶扩展
		s, l := extend(i, i+1, s) 
		if l > maxLen {
			start, maxLen = s, l
		}
	}
	return s[start : start+maxLen]
}

func extend(i, j int, s string) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i, j = i-1, j+1
	}
	return i+1, j-i-1
}
