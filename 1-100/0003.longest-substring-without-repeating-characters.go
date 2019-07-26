/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	hash := make(map[rune]int) // value to index
	maxLen, left := 0, 0
	for right, c := range s {
		if _, ok := hash[c]; ok && hash[c] >= left {
			left = hash[c] + 1
		}
		hash[c] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
