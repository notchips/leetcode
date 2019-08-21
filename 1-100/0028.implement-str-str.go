/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */
 func strStr(haystack string, needle string) int {
	next := getNext(needle)
	i, j := 0, 0
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

func getNext(needle string) []int {
	next := make([]int, len(needle)+1)
	next[0] = -1
	for i, j := 0, -1; i < len(needle); {
		if j == -1 || needle[i] == needle[j] {
			i, j = i+1, j+1
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}
