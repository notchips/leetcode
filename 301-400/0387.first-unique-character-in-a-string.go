/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */
func firstUniqChar(s string) int {
	hash := make([]int, 26)
	for _, c := range s {
		hash[c-'a']++
	}
	for i, c := range s {
		if hash[c-'a'] < 2 {
			return i
		}
	}
	return -1
}
