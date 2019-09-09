/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !isNumber(s[i]) && !isLetter(s[i]) {
			i++
		}
		for i < j && !isNumber(s[j]) && !isLetter(s[j]) {
			j--
		}
		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
	}
	return true
}

func isNumber(b byte) bool {
	return '0' <= b && b <= '9'
}

func isLetter(b byte) bool {
	return ('A' <= b && b <= 'Z') ||
		('a' <= b && b <= 'z')
}

func toLower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + 'a' - 'A'
	}
	return b
}
