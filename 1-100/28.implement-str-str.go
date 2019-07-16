/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */
 // KMP
 func strStr(haystack string, needle string) int {
	i, j := 0, 0
	next := getNext(needle)
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i, j = i+1, j+1
		} else {
			j = next[j]
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}

func getNext(str string) []int {
	next := make([]int, len(str)+1)
	next[0] = -1
	for i, j := 0, -1; i < len(str); {
		if j == -1 || str[i] == str[j] {
			i, j = i+1, j+1
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}
